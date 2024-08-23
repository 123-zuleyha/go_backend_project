package main

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/123-zuleyha/go_backend_project/db"
	"github.com/123-zuleyha/go_backend_project/internal/answer"
	"github.com/123-zuleyha/go_backend_project/internal/auth"
	"github.com/123-zuleyha/go_backend_project/internal/choice"
	"github.com/123-zuleyha/go_backend_project/internal/code"
	codeAnswer "github.com/123-zuleyha/go_backend_project/internal/code-answer"
	codeSubmission "github.com/123-zuleyha/go_backend_project/internal/code-submission"
	"github.com/123-zuleyha/go_backend_project/internal/lesson"
	"github.com/123-zuleyha/go_backend_project/internal/question"
	"github.com/123-zuleyha/go_backend_project/internal/quiz"
	"github.com/123-zuleyha/go_backend_project/internal/types"
	"github.com/123-zuleyha/go_backend_project/internal/user"
	userAnswer "github.com/123-zuleyha/go_backend_project/internal/user-answer"
	userQuiz "github.com/123-zuleyha/go_backend_project/internal/user-quiz"

	"github.com/123-zuleyha/go_backend_project/pkg/middleware"

	rabbitmq "github.com/123-zuleyha/go_backend_project/notification_service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	conn, ch := rabbitmq.Connect()
	defer conn.Close()
	defer ch.Close()

	// Mesaj gönderme
	rabbitmq.PublishMessage(ch, "my_queue", "Hello, RabbitMQ!")

	rabbitmq.ConsumeMessages(ch, "my_queue")

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	go func() {
		select {}
	}()

	db.Connect()

	database := db.DB
	userRepository := user.NewUserRepository(database)
	userService := user.NewUserService(userRepository)
	userHandler := user.NewUserHandler(userService)

	lessonRepository := lesson.NewLessonRepository(database)
	lessonService := lesson.NewLessonService(lessonRepository)
	lessonHandler := lesson.NewLessonHandler(lessonService)

	answerRepository := answer.NewAnswerRepository(database)
	answerService := answer.NewAnswerService(answerRepository)
	answerHandler := answer.NewAnswerHandler(answerService)

	choiceRepository := choice.NewChoiceRepository(database)
	choiceService := choice.NewChoiceService(choiceRepository)
	choiceHandler := choice.NewChoiceHandler(choiceService)

	codeAnswerRepository := codeAnswer.NewCodeAnswerRepository(database)
	codeAnswerService := codeAnswer.NewCodeAnswerService(codeAnswerRepository)
	codeAnswerHandler := codeAnswer.NewCodeAnswerHandler(codeAnswerService)

	codeSubmissionRepository := codeSubmission.NewCodeSubmissionRepository(database)
	codeSubmissionService := codeSubmission.NewCodeSubmissionService(codeSubmissionRepository)
	codeSubmissionHandler := codeSubmission.NewCodeSubmissionHandler(codeSubmissionService)

	codeRepository := code.NewCodeRepository(database)
	codeService := code.NewCodeService(codeRepository)
	codeHandler := code.NewCodeHandler(codeService)

	questionRepository := question.NewQuestionRepository(database)
	questionService := question.NewQuestionService(questionRepository)
	questionHandler := question.NewQuestionHandler(questionService)

	quizRepository := quiz.NewQuizRepository(database)
	quizService := quiz.NewQuizService(quizRepository)
	quizHandler := quiz.NewQuizHandler(quizService)

	userTypeRepository := types.NewUserTypeRepository(database)
	userTypeService := types.NewUserTypeService(userTypeRepository)
	userTypeHandler := types.NewUserTypeHandler(userTypeService)

	userAnswerRepository := userAnswer.NewUserAnswerRepository(database)
	userAnswerService := userAnswer.NewUserAnswerService(userAnswerRepository)
	userAnswerHandler := userAnswer.NewUserAnswerHandler(userAnswerService)

	userQuizRepository := userQuiz.NewUserQuizRepository(database)
	userQuizService := userQuiz.NewUserQuizService(userQuizRepository)
	userQuizHandler := userQuiz.NewUserQuizHandler(userQuizService)

	authRepository := auth.NewAuthRepository(database)
	authService := auth.NewAuthService(authRepository, userRepository)
	authHandler := auth.NewAuthHandler(authService)

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET, POST, PUT, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	})

	api := app.Group("/api/v1")
	api.Post("/login", authHandler.Login)
	api.Post("/register", authHandler.Register)
	api = api.Use(middleware.AuthMiddleware)
	userApi := api.Group("/users")
	userApi.Get("/", userHandler.GetUsers)
	userApi.Get("/:id", userHandler.GetUserByID)
	userApi.Post("/", userHandler.CreateUser)
	userApi.Put("/:id", userHandler.UpdateUser)
	userApi.Delete("/:id", userHandler.DeleteUser)
	userApi.Get("/email/:email", userHandler.GetUserByEmail)
	userApi.Get("/name/:username", userHandler.GetUserByUsername)

	studentAPI := api.Group("/students")
	studentAPI.Get("/", userHandler.GetStudents)
	studentAPI.Get("/:lessonID", lessonHandler.GetStudentsByLesson)
	studentAPI.Get("/:lessonID/:userID", userHandler.AddLessonToUser)

	studentLessonAPI := api.Group("/student-lesson")
	studentLessonAPI.Get("/:lessonID", lessonHandler.GetStudentsByNotInLesson)
	studentLessonAPI.Get("/remove/:lessonID/:userID", userHandler.RemoveLessonFromUser)

	studentQuizAPI := api.Group("/student-quiz")
	studentQuizAPI.Get("/:userID", userQuizHandler.GetUsersQuizByUserID)
	studentQuizAPI.Get("/lesson/:quizID", userHandler.GetUsersQuizzesByQuizID)

	studentAnswerAPI := api.Group("/student-answer")
	studentAnswerAPI.Get("/:userID/:quizID", userQuizHandler.GetUserQuizWithAnswersByUserAndQuizID)

	studentTeacherAPI := api.Group("/student-teacher")
	studentTeacherAPI.Get("/:lessonID", userHandler.GetStudentsByTeacher)

	teacherAPI := api.Group("/teacher")
	teacherAPI.Get("/:userID", userHandler.SetTeacher)
	teacherAPI.Get("/", userHandler.GetTeacher)

	teacherLessonAPI := api.Group("/teacher-lesson")
	teacherLessonAPI.Get("/:userID", lessonHandler.GetLessonByTeacher)

	lessonAPI := api.Group("/lessons")
	lessonAPI.Get("/", lessonHandler.GetLessons)
	lessonAPI.Get("/:id", lessonHandler.GetLessonByID)
	lessonAPI.Post("/", lessonHandler.CreateLessons)
	lessonAPI.Delete("/:id", lessonHandler.DeleteLessons)
	lessonAPI.Put("/", lessonHandler.UpdateLessons)

	answerAPI := api.Group("/answers")
	answerAPI.Get("/", answerHandler.GetAnswers)
	answerAPI.Get("/:id", answerHandler.GetAnswerByID)
	answerAPI.Post("/", answerHandler.CreateAnswer)
	answerAPI.Put("/", answerHandler.UpdateAnswer)
	answerAPI.Delete("/", answerHandler.DeleteAnswer)

	answerQuestionAPI := api.Group("/answer-question")
	answerQuestionAPI.Get("/", answerHandler.GetAnswersByQuestionID)

	choiceAPI := api.Group("/choices")
	choiceAPI.Get("/", choiceHandler.GetChoices)
	choiceAPI.Get("/:id", choiceHandler.GetChoiceByID)
	choiceAPI.Post("/", choiceHandler.CreateChoices)
	choiceAPI.Put("/", choiceHandler.UpdateChoices)
	choiceAPI.Delete("/:id", choiceHandler.DeleteChoices)

	questionChoiceAPI := api.Group("/question-choices")
	questionChoiceAPI.Get("/:questionID", choiceHandler.GetChoicesByQuestionID)
	questionChoiceAPI.Get("quiz/:quizID", choiceHandler.GetQuestionsWithChoicesByQuizID)

	codeAnswerAPI := api.Group("/code-answers")
	codeAnswerAPI.Get("/", codeAnswerHandler.GetCodeAnswers)
	codeAnswerAPI.Get("/:id", codeAnswerHandler.GetCodeAnswerByID)
	codeAnswerAPI.Post("/", codeAnswerHandler.CreateCodeAnswer)
	codeAnswerAPI.Delete("/:id", codeAnswerHandler.DeleteCodeAnswer)
	codeAnswerAPI.Put("/", codeAnswerHandler.UpdateCodeAnswer)

	codeAnswerUserAPI := api.Group("/code-answer-users")
	codeAnswerUserAPI.Get("/:user_id", codeAnswerHandler.GetCodeAnswersByUserID)

	codeSubmissionAPI := api.Group("/code-submissions")
	codeSubmissionAPI.Get("/", codeSubmissionHandler.GetCodeSubmissions)
	codeSubmissionAPI.Get("/:id", codeSubmissionHandler.GetCodeSubmissionByID)
	codeSubmissionAPI.Post("/", codeSubmissionHandler.CreateCodeSubmission)
	codeSubmissionAPI.Delete("/:id", codeSubmissionHandler.DeleteCodeSubmission)
	codeSubmissionAPI.Put("/", codeSubmissionHandler.UpdateCodeSubmission)

	codeSubmissionCodeAPI := api.Group("/code-submission-code")
	codeSubmissionCodeAPI.Get("/:codeID", codeSubmissionHandler.GetCodeSubmissionsByCodeID)

	codeAPI := api.Group("/codes")
	codeAPI.Get("/", codeHandler.GetCodes)
	codeAPI.Get("/:id", codeHandler.GetCodeByID)
	codeAPI.Post("/", codeHandler.CreateCode)
	codeAPI.Put("/", codeHandler.UpdateCode)
	codeAPI.Delete("/:id", codeHandler.DeleteCode)

	codeLessonAPI := api.Group("/code-lesson")
	codeLessonAPI.Get("/:lessonID", codeHandler.GetCodesByLessonID)

	codeTeacherAPI := api.Group("/code-teacher")
	codeTeacherAPI.Get("/:userID", codeHandler.GetCodesByTeacherID)

	codeUserAPI := api.Group("/code-user")
	codeUserAPI.Get("/:userID", codeHandler.GetUsersCodes)

	questionAPI := api.Group("/questions")
	questionAPI.Get("/", questionHandler.GetQuestions)
	questionAPI.Get("/:id", questionHandler.GetQuestionByID)
	questionAPI.Post("/", questionHandler.CreateQuestions)
	questionAPI.Put("/", questionHandler.UpdateQuestion)
	questionAPI.Delete("/:id", questionHandler.DeleteQuestion)
	questionAPI.Get("/quiz/:id", questionHandler.GetQuestionsByQuizID)

	quizAPI := api.Group("/quizzes")
	quizAPI.Get("/", quizHandler.GetQuizzes)
	quizAPI.Post("/", quizHandler.CreateQuizzes)
	quizAPI.Put("/", quizHandler.UpdateQuiz)
	quizAPI.Delete("/:id", quizHandler.DeleteQuiz)
	quizAPI.Get("/:id", quizHandler.GetQuizByID)

	teacherQuizAPI := api.Group("/teacher-quizzes")
	teacherQuizAPI.Get("/:teacher_id", quizHandler.GetQuizByTeacher)

	userTypeAPI := api.Group("/user-types")
	userTypeAPI.Get("/", userTypeHandler.GetUserTypes)
	userTypeAPI.Post("/", userTypeHandler.CreateUserType)

	userAnswerAPI := api.Group("/user-answers")
	userAnswerAPI.Get("/", userAnswerHandler.GetUserAnswers)
	userAnswerAPI.Get("/:id", userAnswerHandler.GetUserAnswerByID)
	userAnswerAPI.Post("/", userAnswerHandler.CreateUserAnswers)
	userAnswerAPI.Delete("/:id", userAnswerHandler.DeleteUserAnswer)
	userAnswerAPI.Put("/", userAnswerHandler.UpdateUserAnswer)

	userAnswerQuizAPI := api.Group("/user-answer-quizzes")
	userAnswerQuizAPI.Get("/:questionID", userAnswerHandler.GetUserAnswerByQuestionID)

	userQuizAPI := api.Group("/user-quizzes")
	userQuizAPI.Get("/", userQuizHandler.GetUserQuizzes)
	userQuizAPI.Get("/:id", userQuizHandler.GetUserQuizByID)
	userQuizAPI.Post("/", userQuizHandler.CreateUserQuiz)
	userQuizAPI.Put("/", userQuizHandler.UpdateUserQuiz)
	userQuizAPI.Delete("/:id", userQuizHandler.DeleteUserQuiz)
	userQuizAPI.Get("/:quizID/:userID", userQuizHandler.GetUserQuizByUserAndQuizID)
	userQuizAPI.Get("/joined/:quizID", codeHandler.GetUsersCodes)

	app.Listen(":8081")
}
