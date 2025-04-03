package usecase

import (
	"encoding/json"
	"io"
	"net/http"
	"sort"
	"strconv"
	"testtesttest/internal/c_ud"
	"testtesttest/pkg/logger"
	"testtesttest/pkg/postgres"
	simplehttp "testtesttest/pkg/simple_http"

	"github.com/gin-gonic/gin"
)

type CUD struct {
	DB        *postgres.Postgres
	Logger    logger.Logger
	GenderApi string
	AgeApi    string
	NationApi string
}

func NewAdd(DB *postgres.Postgres, Logger logger.Logger, GenderApi string, AgeApi string, NationApi string) c_ud.C_UDUC {
	Logger.Info("Create, Update, Delete service ready")
	return &CUD{DB: DB, Logger: Logger, AgeApi: AgeApi, GenderApi: GenderApi, NationApi: NationApi}
}

const AddQuery = "insert into Users (Name, Surname, Patronymic, Nation, Gender, Age) values ($1, $2, $3, $4, $5, $6) returning ID"
const DelQuery = "delete from Users where ID = $1;"

func (CUD *CUD) AddUser(c *gin.Context) {
	var ID int
	var GeneralInfo c_ud.UserInfo
	var AgeInfo c_ud.AgeRequest
	var GenderInfo c_ud.GenderRequst
	var NationInfo c_ud.NationRequest
	db := CUD.DB.Get()

	CUD.Logger.Debug("Got Add request")

	if err := c.ShouldBindBodyWithJSON(&GeneralInfo); err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if GeneralInfo.Name == "" || GeneralInfo.Surname == "" {
		CUD.Logger.Error("Name or surname absent")
		c.JSON(http.StatusBadRequest, "Name or surname absent")
		return
	}
	data, err := simplehttp.MakeRequest(CUD.AgeApi, GeneralInfo.Name)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(data, &AgeInfo)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if AgeInfo.Age == 0 {
		//		CUD.Logger.Error("Age for " + GeneralInfo.Name + " not found")
		c.JSON(http.StatusBadRequest, "Age for "+GeneralInfo.Name+" not found")
		return
	}
	data, err = simplehttp.MakeRequest(CUD.GenderApi, GeneralInfo.Name)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(data, &GenderInfo)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if GenderInfo.Gender == "" {
		//	CUD.Logger.Error("Gender for " + GeneralInfo.Name + " not found")
		c.JSON(http.StatusBadRequest, "Gender for "+GeneralInfo.Name+"not found")
		return
	}
	data, err = simplehttp.MakeRequest(CUD.NationApi, GeneralInfo.Name)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = json.Unmarshal(data, &NationInfo)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if NationInfo.Country == nil {
		c.JSON(http.StatusBadRequest, "Nation for "+GeneralInfo.Name+"not found")
		CUD.Logger.Error("Nation for " + GeneralInfo.Name + "not found")
		return
	}
	GeneralInfo.Age = AgeInfo.Age
	GeneralInfo.Gender = GenderInfo.Gender
	GeneralInfo.Nation = sortNation(NationInfo)

	db.QueryRow(AddQuery, GeneralInfo.Name, GeneralInfo.Surname, GeneralInfo.Patronymic, GeneralInfo.Nation, GeneralInfo.Gender, GeneralInfo.Age).Scan(&ID)
	CUD.Logger.Debug("User added")
	c.JSON(http.StatusOK, ID)
}

func (CUD *CUD) DeleteUser(c *gin.Context) {
	CUD.Logger.Debug("Got Delete request")
	IDString := c.Query("ID")
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	db := CUD.DB.Get()
	res, err := db.Exec(DelQuery, ID)
	if err != nil || res == nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		CUD.Logger.Error(err)
		return
	}
	if rows == 0 {
		c.JSON(http.StatusNoContent, "No entry found")
		//		CUD.Logger.Error("no entry")
		return
	}
	CUD.Logger.Debug("Entry Deleted")
	c.JSON(http.StatusOK, "Entry Deleted")
}

func (CUD *CUD) ChangeUser(c *gin.Context) {
	var Values map[string]interface{}
	var iter = 1
	var Query []interface{}
	QueryStr := "UPDATE Users SET "
	CUD.Logger.Debug("Got Change request")
	IDString := c.Query("ID")
	ID, err := strconv.Atoi(IDString)
	if err != nil {
		CUD.Logger.Error(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
	data, err := io.ReadAll(c.Request.Body)
	json.Unmarshal(data, &Values)
	if val, ok := Values["name"]; ok {
		Query = append(Query, val)
		QueryStr = QueryStr + "Name" + "=$" + strconv.Itoa(iter) + ","
		iter++
	}
	if val, ok := Values["surname"]; ok {
		Query = append(Query, val)
		QueryStr = QueryStr + "Surname" + "=$" + strconv.Itoa(iter) + ","
		iter++
	}
	if val, ok := Values["patronymic"]; ok {
		Query = append(Query, val)
		QueryStr = QueryStr + "Patronymic" + "=$" + strconv.Itoa(iter) + ","
		iter++
	}
	if val, ok := Values["nation"]; ok {
		Query = append(Query, val)
		QueryStr = QueryStr + "Nation" + "=$" + strconv.Itoa(iter) + ","
		iter++
	}
	if val, ok := Values["age"]; ok {
		Query = append(Query, val)
		QueryStr = QueryStr + "Age" + "=$" + strconv.Itoa(iter) + ","
		iter++
	}
	if val, ok := Values["gender"]; ok {
		Query = append(Query, val)
		QueryStr = QueryStr + "Gender" + "=$" + strconv.Itoa(iter) + " "
		iter++
	}
	QueryStr = QueryStr[:len(QueryStr)-1]
	QueryStr = QueryStr + " WHERE ID=" + strconv.Itoa(ID) + " returning ID;"
	if len(Values) > iter-1 || iter == 1 {
		CUD.Logger.Debug("Bad params or no params")
		c.JSON(http.StatusBadRequest, "Bad params or no params")
		return
	}
	db := CUD.DB.Get()
	db.QueryRow(QueryStr, Query...).Scan(&ID)
	FinalID := strconv.Itoa(ID)
	CUD.Logger.Debug("Query str ready " + QueryStr)
	c.JSON(http.StatusOK, "Entry modifed ID:"+FinalID)
}

// Yes, it is sorted, but we cant be sure, so lets sort it

func sortNation(Nation c_ud.NationRequest) string {
	sort.SliceIsSorted(Nation.Country, func(i, j int) bool {
		return Nation.Country[i].Probability < Nation.Country[j].Probability
	})
	return Nation.Country[0].CountryID
}
