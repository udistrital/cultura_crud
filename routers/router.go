// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	//"github.com/udistrital/cultura_crud/controllers"
	"cultura_crud/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/actividad_cultural",
			beego.NSInclude(
				&controllers.ActividadCulturalController{},
			),
		),

		beego.NSNamespace("/actividad_grupo_cultural",
			beego.NSInclude(
				&controllers.ActividadGrupoCulturalController{},
			),
		),

		beego.NSNamespace("/grupo_cultural",
			beego.NSInclude(
				&controllers.GrupoCulturalController{},
			),
		),

		beego.NSNamespace("/evidencia_actividad_cultural",
			beego.NSInclude(
				&controllers.EvidenciaActividadCulturalController{},
			),
		),

		beego.NSNamespace("/horarios_grupo_cultural",
			beego.NSInclude(
				&controllers.HorariosGrupoCulturalController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
