package api

import (
	v1 "github.com/asilbek17071/apigateway/api/handler/v1"
	"github.com/asilbek17071/apigateway/config"
	"github.com/asilbek17071/apigateway/middlewares"
	"github.com/asilbek17071/apigateway/pkg/logger"
	"github.com/asilbek17071/apigateway/services"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Conf           config.Config
	Logger         logger.Logger
	ServiceManager services.IServiceManager
}

func New(option Option) *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middlewares.Cors())

	handlerV1 := v1.New(&v1.HandlerV1Config{
		Logger:         option.Logger,
		ServiceManager: option.ServiceManager,
		Cfg:            option.Conf,
	})

	planner := router.Group("/Planner")
	planner.POST("/user/create/", handlerV1.UserCreate)
	planner.POST("/user/login/", handlerV1.UserLogin)
	planner.GET("/user/byid/", handlerV1.UserGet)
	planner.GET("/user/byname/", handlerV1.UserByName)
	planner.GET("/user/list/", handlerV1.UserList)
	planner.PUT("/user/update/", handlerV1.UserUpdate)
	planner.DELETE("/user/delete/", handlerV1.UserDelete)
	planner.GET("/signup/user/", handlerV1.UserSignUp)

	planner.GET("/dashboard/lead/", handlerV1.LeadActive)
	planner.GET("/dashboard/student/", handlerV1.StudentActive)
	planner.GET("/dashboard/toplist/", handlerV1.StudentTopList)
	planner.GET("/dashboard/deletelist/", handlerV1.StudentDeleteList)
	planner.GET("/dashboard/student_profit/", handlerV1.StudentProfit)
	planner.GET("/dashboard/profit/", handlerV1.Profit)

	planner.POST("/teacher/create/", handlerV1.TeacherCreate)
	planner.POST("/teacher/login/", handlerV1.TeacherLogin)
	planner.GET("/teacher/byid/", handlerV1.TeacherGet)
	planner.GET("/teacher/byname/", handlerV1.TeacherByName)
	planner.GET("/signup/teacher/", handlerV1.TeacherSignUp)
	planner.GET("/teacher/list/", handlerV1.TeacherList)
	planner.PUT("/teacher/update/", handlerV1.TeacherUpdate)
	planner.DELETE("/teacher/delete/", handlerV1.TeacherDelete)

	planner.POST("/student/create/", handlerV1.StudentCreate)
	planner.POST("/student/phonecreate/", handlerV1.StudentPhoneCreate)
	planner.POST("/student/login/", handlerV1.StudentLogin)
	planner.GET("/student/byid/", handlerV1.StudentGet)
	planner.GET("/signup/student/", handlerV1.StudentSignUp)
	planner.GET("/student/listphone/", handlerV1.StudentPhoneGet)
	planner.GET("/student/list/", handlerV1.StudentList)
	planner.GET("/student/grouplist/", handlerV1.StudentGroupList)
	planner.PUT("/student/update/", handlerV1.StudentUpdate)
	planner.DELETE("/student/delete/", handlerV1.StudentDelete)

	planner.POST("/direction/create/", handlerV1.DirectionCreate)
	planner.GET("/direction/byid/", handlerV1.DirectionGet)
	planner.GET("/direction/list/", handlerV1.DirectionList)
	planner.PUT("/direction/update/", handlerV1.DirectionUpdate)
	planner.DELETE("/direction/delete/", handlerV1.DirectionDelete)

	planner.POST("/personal/create/", handlerV1.PersonalCreate)
	planner.GET("/personal/byid/", handlerV1.PersonalGet)
	planner.GET("/personal/position/", handlerV1.Position)
	planner.GET("/personal/list/", handlerV1.PersonalList)
	planner.GET("/personals/", handlerV1.Personals)
	planner.PUT("/personal/update/", handlerV1.PersonalUpdate)
	planner.DELETE("/personal/delete/", handlerV1.PersonalDelete)

	planner.POST("/type/spending/create/", handlerV1.SpendingTypeCreate)
	planner.GET("/type/spending/byid/", handlerV1.SpendingTypeGet)
	planner.GET("/type/spending/list/", handlerV1.SpendingTypeList)
	planner.PUT("/type/spending/update/", handlerV1.SpendingTypeUpdate)
	planner.DELETE("/type/spending/delete/", handlerV1.SpendingTypeDelete)

	planner.POST("/type/salary/create/", handlerV1.SalaryTypeCreate)
	planner.GET("/type/salary/byid/", handlerV1.SalaryTypeGet)
	planner.GET("/type/salary/list/", handlerV1.SalaryTypeList)
	planner.PUT("/type/salary/update/", handlerV1.SalaryTypeUpdate)
	planner.DELETE("/type/salary/delete/", handlerV1.SalaryTypeDelete)

	planner.POST("/salary/teacher/create/", handlerV1.TeacherSalaryCreate)
	planner.GET("/salary/teacher/byid/", handlerV1.TeacherSalaryGet)
	planner.GET("/salary/teacher/list/", handlerV1.TeacherSalaryList)
	planner.PUT("/salary/teacher/update/", handlerV1.TeacherSalaryUpdate)
	planner.DELETE("/salary/teacher/delete/", handlerV1.TeacherSalaryDelete)

	planner.POST("/salary/admin/create/", handlerV1.UserSalaryCreate)
	planner.GET("/salary/admin/byid/", handlerV1.UserSalaryGet)
	planner.GET("/salary/admin/list/", handlerV1.UserSalaryList)
	planner.PUT("/salary/admin/update/", handlerV1.UserSalaryUpdate)
	planner.DELETE("/salary/admin/delete/", handlerV1.UserSalaryDelete)

	planner.POST("/group/create/", handlerV1.GroupCreate)
	planner.GET("/group/byid/", handlerV1.GroupGet)
	planner.GET("/group/list/", handlerV1.GroupList)
	planner.GET("/group/byteacherid/list/", handlerV1.TeacherGroup)
	planner.GET("/group/withlist/", handlerV1.GroupWithList)
	planner.GET("/group/roomlist/", handlerV1.GroupRoomList)
	planner.GET("/group/attendancelist/", handlerV1.GroupAttendanceList)
	planner.PUT("/group/update/", handlerV1.GroupUpdate)
	planner.DELETE("/group/delete/", handlerV1.GroupDelete)

	planner.POST("/branch/create/", handlerV1.BranchCreate)
	planner.GET("/branch/byid/", handlerV1.BranchGet)
	planner.GET("/branch/withroom/", handlerV1.BranchWithRooms)
	planner.GET("/branch/list/", handlerV1.BranchList)
	planner.PUT("/branch/update/", handlerV1.BranchUpdate)
	planner.DELETE("/branch/delete/", handlerV1.BranchDelete)

	planner.POST("/room/create/", handlerV1.RoomCreate)
	planner.GET("/room/byid/", handlerV1.RoomGet)
	planner.GET("/room/list/", handlerV1.RoomList)
	planner.PUT("/room/update/", handlerV1.RoomUpdate)
	planner.DELETE("/room/delete/", handlerV1.RoomDelete)

	planner.POST("/finance/salary/create/", handlerV1.SalaryFinanceCreate)
	planner.GET("/finance/salary/", handlerV1.SalaryFinanceGet)

	planner.POST("/payment_system/create/", handlerV1.PySysCreate)
	planner.GET("/payment_system/byid/", handlerV1.PySysGet)
	planner.GET("/payment_system/list/", handlerV1.PySysList)
	planner.PUT("/payment_system/update/", handlerV1.PySysUpdate)
	planner.DELETE("/payment_system/delete/", handlerV1.PySysDelete)

	planner.POST("/attendance/create/", handlerV1.AttendanceCreate)
	planner.GET("/attendance/byid/", handlerV1.AttendanceGet)
	planner.GET("/attendance/list/", handlerV1.AttendanceList)
	planner.GET("/attendance/list/participate/", handlerV1.ParticipateResp)
	planner.PUT("/attendance/update/", handlerV1.AttendanceUpdate)
	planner.DELETE("/attendance/delete/", handlerV1.AttendanceDelete)

	planner.POST("/lead/create/", handlerV1.LeadCreate)
	planner.POST("/lead/phonecreate/", handlerV1.LeadPhoneCreate)
	planner.GET("/lead/byid/", handlerV1.LeadGet)
	planner.GET("/lead/listphone/", handlerV1.LeadPhoneGet)
	planner.GET("/lead/list/", handlerV1.LeadList)
	planner.PUT("/lead/update/", handlerV1.LeadUpdate)
	planner.DELETE("/lead/delete/", handlerV1.LeadDelete)

	planner.POST("/payment/create/", handlerV1.PaymentCreate)
	planner.GET("/payment/byid/", handlerV1.PaymentGet)
	planner.GET("/payment/list/", handlerV1.PaymentList)
	planner.GET("/payment/search/", handlerV1.PaymentSearch)
	planner.GET("/payment/searchlist/", handlerV1.PaymentSearchList)
	planner.PUT("/payment/update/", handlerV1.PaymentUpdate)
	planner.DELETE("/payment/delete/", handlerV1.PaymentUpdate)

	planner.POST("/spending/create/", handlerV1.SpendingCreate)
	planner.GET("/spending/byid/", handlerV1.SpendingGet)
	planner.GET("/spending/list/", handlerV1.SpendingList)
	planner.PUT("/spending/update/", handlerV1.SpendingUpdate)
	planner.DELETE("/spending/delete/", handlerV1.SpendingDelete)

	planner.POST("/task/create/", handlerV1.TaskCreate)
	planner.GET("/task/byid/", handlerV1.TaskGet)
	planner.GET("/task/list/", handlerV1.TaskList)
	planner.PUT("/task/update/", handlerV1.TaskUpdate)
	planner.DELETE("/task/delete/", handlerV1.TaskDelete)

	planner.POST("/sale/create/", handlerV1.SaleCreate)
	planner.GET("/sale/byid/", handlerV1.SaleGet)
	planner.GET("/sale/list/", handlerV1.SaleList)
	planner.PUT("/sale/update/", handlerV1.SaleUpdate)
	planner.DELETE("/sale/delete/", handlerV1.SaleDelete)

	planner.POST("/defmsg/create/", handlerV1.DefMsgCreate)
	planner.GET("/defmsg/byid/", handlerV1.DefMsgGet)
	planner.GET("/defmsg/list/", handlerV1.DefMsgList)
	planner.PUT("/defmsg/update/", handlerV1.DefMsgUpdate)
	planner.DELETE("/defmsg/delete/", handlerV1.DefMsgDelete)

	planner.POST("/ball/create/", handlerV1.BallCreate)
	planner.GET("/ball/list/", handlerV1.BallList)
	planner.PUT("/ball/update/", handlerV1.BallUpdate)
	planner.DELETE("/ball/delete/", handlerV1.BallDelete)

	return router
}
