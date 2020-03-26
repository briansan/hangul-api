package main

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func handlePing(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func handlePronounce(c echo.Context) error {
	ch := c.QueryParam("ch")
	logrus.WithField("ch", ch).Info("+pronounce")

	val := url.Values{}
	val.Add("msg", ch)
	val.Add("lang", "Seoyeon")
	val.Add("source", "ttsmp3")

	resp, err := http.DefaultClient.PostForm("https://ttsmp3.com/makemp3_new.php", val)
	if err != nil {
		logrus.WithError(err).Error("pronounce: failed to make request")
		return c.JSON(http.StatusInternalServerError, nil)
	}
	defer resp.Body.Close()

	obj := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		logrus.WithError(err).Error("pronounce: failed to make request")
		return c.JSON(http.StatusInternalServerError, nil)
	}

	logrus.WithField("resp", obj).Info("-pronounce")
	return c.JSON(http.StatusOK, obj)
}

func main() {
	e := echo.New()
	e.GET("/pronounce", handlePronounce)
	e.GET("/", handlePing)

	logrus.Info("Starting server")
	e.Start(":5250")
}
