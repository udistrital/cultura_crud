package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadCulturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:ActividadGrupoCulturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:EvidenciaActividadCulturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:GrupoCulturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"],
        beego.ControllerComments{
            Method: "Post",
            Router: "/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: "/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"],
        beego.ControllerComments{
            Method: "Put",
            Router: "/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"] = append(beego.GlobalControllerRouter["github.com/udistrital/cultura_crud/controllers:HorariosGrupoCulturalController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: "/:id",
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
