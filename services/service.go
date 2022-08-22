package services

import (
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"

	"github.com/asilbek17071/apigateway/config"
	pbCourse "github.com/asilbek17071/apigateway/genproto/course_service"
	pbFinance "github.com/asilbek17071/apigateway/genproto/finance_service"
	pbStudent "github.com/asilbek17071/apigateway/genproto/student_service"
	pbSystem "github.com/asilbek17071/apigateway/genproto/system_service"
	pbTeacher "github.com/asilbek17071/apigateway/genproto/teacher_service"
	pbUser "github.com/asilbek17071/apigateway/genproto/user_service"
)

type IServiceManager interface {
	UserService() pbUser.UserServiceClient
	TeacherService() pbTeacher.TeacherServiceClient
	StudentService() pbStudent.StudentServiceClient
	SystemService() pbSystem.SystemServiceClient
	FinanceService() pbFinance.FinanceServiceClient
	CourseService() pbCourse.CourseServiceClient
}

type serviceManager struct {
	userService    pbUser.UserServiceClient
	teacherService pbTeacher.TeacherServiceClient
	studentService pbStudent.StudentServiceClient
	systemService  pbSystem.SystemServiceClient
	financeService pbFinance.FinanceServiceClient
	courseService  pbCourse.CourseServiceClient
}

func (s *serviceManager) UserService() pbUser.UserServiceClient {
	return s.userService
}

func (s *serviceManager) TeacherService() pbTeacher.TeacherServiceClient {
	return s.teacherService
}

func (s *serviceManager) StudentService() pbStudent.StudentServiceClient {
	return s.studentService
}

func (s *serviceManager) SystemService() pbSystem.SystemServiceClient {
	return s.systemService
}

func (s *serviceManager) FinanceService() pbFinance.FinanceServiceClient {
	return s.financeService
}

func (s *serviceManager) CourseService() pbCourse.CourseServiceClient {
	return s.courseService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")

	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connTeacher, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.TeacherServiceHost, conf.TeacherServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connStudent, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.StudentServiceHost, conf.StudentServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connSystem, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.SystemServiceHost, conf.SystemServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connFinance, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.FinanceServiceHost, conf.FinanceServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	connCourse, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CourseServiceHost, conf.CourseServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		userService:    pbUser.NewUserServiceClient(connUser),
		teacherService: pbTeacher.NewTeacherServiceClient(connTeacher),
		studentService: pbStudent.NewStudentServiceClient(connStudent),
		systemService:  pbSystem.NewSystemServiceClient(connSystem),
		financeService: pbFinance.NewFinanceServiceClient(connFinance),
		courseService:  pbCourse.NewCourseServiceClient(connCourse),
	}

	return serviceManager, nil
}
