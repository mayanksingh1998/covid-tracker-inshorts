// Package api cases API.
//
// the purpose of this application is to provide the data of
// covid cases in a particular city
//
//
// Documentation of cases api
//
//     Schemes: http
//     Host: localhost
//     BasePath: /cases/
//     Version: 1.1
//     Contact: Mayank <mayank.singh@gmail.com>
//
//     Consumes:
//     - query
//
//     Produces:
//     - application/json
//
// swagger:meta
package api

import (
	"covid-tracker/contract"
	"covid-tracker/service"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"reflect"
)



func  CasesHandler(c echo.Context) (err error) {
	//var VideoAPIResponse videoAPIResponse
	//var video *contract.Video
	u := new(contract.UserLocation)
	if err := c.Bind(u); err != nil {
		return nil
	}
	fmt.Println(reflect.TypeOf(u.Longitude))
	fmt.Println(u.Longitude)
	state := service.GetUserLocationCases(u)
	//user := UserDTO{
	//	Name: u.Name,
	//	Email: u.Email,
	//	IsAdmin: false // because you could accidentally expose fields that should not be bind
	//}
	//sourceIDQ, _ := c.GetQuery("source_id")
	//sourceID, sourceIDUUIDParseErr := uuid.Parse(sourceIDQ)
	//var videoID = c.Param("video_id")
	//
	//if sourceIDUUIDParseErr != nil {
	//	var videoId, videoIDUUIDParseErr = uuid.Parse(videoID)
	//	if videoIDUUIDParseErr == nil {
	//		video = API.CourseService.GetVideoByID(&videoId)
	//	}
	//} else{
	//	video = API.CourseService.GetVideoIDBySourceIDAndSourceVideoID(&sourceID, &videoID)
	//}
	//if video == nil {
	//	utils.SetNotFound(c, err.ErrVideoNotFound, gin.H{})
	//	return
	//}
	//
	//VideoAPIResponse.Video = video
	return c.JSON(http.StatusOK, state )
	//utils.SetSuccess(c, VideoAPIResponse)
	//return c.JSON()
}
