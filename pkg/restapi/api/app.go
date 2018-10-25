package api

import (
	"alert-kubesphere-plugin/pkg/client"
	"alert-kubesphere-plugin/pkg/models"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"github.com/emicklei/go-restful/log"
	"github.com/go-openapi/spec"
	"net/http"
)

func (u MonitorResource) senderAlertConfig(request *restful.Request, response *restful.Response) {
	client.SenderAlertConfig(request, response)
}

func (u MonitorResource) sayBye(request *restful.Request, response *restful.Response) {
	name := request.PathParameter("name")
	response.WriteAsJson("bye, " + name)

}

type MonitorResource struct {
}

func Run() {
	u := MonitorResource{}
	restful.DefaultContainer.Add(u.WebService())
	handleSwagger()
	enableCORS()

	log.Printf("Get the API using http://localhost:8080/apidocs.json")
	log.Printf("Open Swagger UI using http://localhost:8080/apidocs/") // ?url=http://localhost:8080/apidocs.json
	log.Print(http.ListenAndServe(":8080", nil))
}

func enableCORS() {
	// Optionally, you may need to enable CORS for the UI to work.
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
		CookiesAllowed: false,
		Container:      restful.DefaultContainer}
	restful.DefaultContainer.Filter(cors.Filter)
}

func handleSwagger() {
	config := restfulspec.Config{
		WebServices: restful.RegisteredWebServices(), // you control what services are visible
		APIPath:     "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(config))
	// Open http://localhost:8080/apidocs/?url=http://localhost:8080/apidocs.json
	http.Handle("/apidocs/", http.StripPrefix("/apidocs/", http.FileServer(http.Dir("C:/Users/Carman/go/src/alert-kubesphere-plugin/swagger-ui/dist"))))
}

func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title: "kubesphere alertmanager restful apis",
			Contact: &spec.ContactInfo{
				Name:  "carman",
				Email: "carmanzhang@yunify.com",
				URL:   "",
			},
			License: &spec.License{
				Name: "MIT License",
				URL:  "http://mit.org",
			},
			Version: "1.0.0",
		},
	}
}

func (u MonitorResource) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/v1/monitoring").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	tags := []string{"monitoring apis"}

	ws.Route(ws.POST("/alert").To(u.senderAlertConfig).
		// user id, user email, resource_type, resource_name, metric_name, condition
		Doc("send alert configmap").
		Operation("PutUserRequest").
		Reads(models.UserRequest{}).
		//Param(myws.BodyParameter("user_info", "user request information").DataType("UserRequest").Required(true)).
		Metadata(restfulspec.KeyOpenAPITags, tags)).
		Produces(restful.MIME_JSON).
		Consumes(restful.MIME_JSON)

	ws.Route(ws.GET("/bye/{name}").To(u.sayBye).
		Doc("test01").
		Param(ws.PathParameter("name", "your name").DataType("string").Required(true).DefaultValue("carman")).
		Metadata(restfulspec.KeyOpenAPITags, tags)).
		Produces(restful.MIME_JSON)

	ws.Route(ws.PUT("/{user-id}").To(u.updateUser).
		// docs
		Doc("update a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(User{})) // from the request
	return ws
}

//type User struct {
//	ID   string `json:"id" description:"identifier of the user"`
//	Name string `json:"name" description:"name of the user" default:"john"`
//	Age  int    `json:"age" description:"age of the user" default:"21"`
//}
