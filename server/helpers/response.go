package helpers

import "github.com/gofiber/fiber/v2"

type ResponseWithData struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
	Data    any    `json:"data"`
}

type ResponseWithoutData struct {
	Status  string `json:"status"`
	Message any    `json:"message"`
}

func Response(c *fiber.Ctx, code int, message any, payload interface{}) error {
	c.Set("Content-Type", "application/json")
	c.Status(code)

	var response interface{}
	status := "Succes"

	if code == 200 {
		status = "OK"
	} else if code == 201 {
		status = "Created"
	} else if code == 400 {
		status = "Bad Request"
	} else if code == 401 {
		status = "Unauthorized"
	} else if code == 404 {
		status = "Not Found"
	} else {
		status = "Internal Server Error"
	}

	if payload != nil {
		response = &ResponseWithData{
			Status:  status,
			Message: message,
			Data:    payload,
		}
	} else {
		response = &ResponseWithoutData{
			Status:  status,
			Message: message,
		}
	}

	return c.JSON(response)
}
