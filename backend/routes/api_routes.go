package routes

import (
	"Rp/backend/controllers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
    r := mux.NewRouter()

    // API routes
    api := r.PathPrefix("/api").Subrouter()
    api.HandleFunc("/student-subjects-only/{studentID}", controllers.GetStudentSubjectsOnlyHandler).Methods("GET", "OPTIONS")
    api.HandleFunc("/save-all-subjects", controllers.SaveAllSubjectsHandler).Methods("POST", "OPTIONS")
    api.HandleFunc("/student-subjects/{studentID}", controllers.GetStudentSubjectsHandler).Methods("GET", "OPTIONS")
    api.HandleFunc("/calculate", controllers.CalculateHandler).Methods("POST", "OPTIONS")
    api.HandleFunc("/calculate-db", controllers.CalculateFromDBHandler).Methods("POST", "OPTIONS")
    api.HandleFunc("/subjects", controllers.GetSubjectsHandler).Methods("GET", "OPTIONS")
    api.HandleFunc("/subject-points/update", controllers.UpdateSubjectPointsHandler).Methods("POST", "OPTIONS")
    api.HandleFunc("/subject-points/{studentID}", controllers.GetSubjectPointsHandler).Methods("GET", "OPTIONS")
    api.HandleFunc("/save-subject-points", controllers.SaveSubjectPointsHandler).Methods("POST", "OPTIONS") 
    api.HandleFunc("/calculate-subject", controllers.CalculateSubjectMarksHandler).Methods("POST", "OPTIONS")
    api.HandleFunc("/student-rewards/{studentID}", controllers.GetStudentRewardsHandler).Methods("GET", "OPTIONS")
    api.HandleFunc("/google", controllers.GoogleLoginHandler).Methods("POST", "OPTIONS")
    
    // Admin routes
    admin := r.PathPrefix("/admin/api").Subrouter()
    admin.HandleFunc("/config", controllers.GetConfigHandler).Methods("GET", "OPTIONS")
    admin.HandleFunc("/config/all", controllers.GetConfigHandler).Methods("GET", "OPTIONS")
    admin.HandleFunc("/config", controllers.UpdateConfigHandler).Methods("POST", "OPTIONS")
    admin.HandleFunc("/departments", controllers.GetDepartmentsHandler).Methods("GET", "OPTIONS")
    admin.HandleFunc("/departments/add", controllers.AddDepartmentHandler).Methods("POST", "OPTIONS")
    admin.HandleFunc("/departments/delete", controllers.DeleteDepartmentHandler).Methods("POST", "OPTIONS")
    admin.HandleFunc("/unfreeze-student", controllers.UnfreezeStudentHandler).Methods("POST", "OPTIONS")
    // Redemption dates routes - moved to admin subrouter
    admin.HandleFunc("/redemption-dates", controllers.GetRedemptionDatesHandler).Methods("GET", "OPTIONS")
    admin.HandleFunc("/redemption-dates", controllers.SaveRedemptionDatesHandler).Methods("POST", "OPTIONS")
    admin.HandleFunc("/redemption-dates/extend", controllers.ExtendRedemptionDeadlineHandler).Methods("POST", "OPTIONS")   

    
    return r
}