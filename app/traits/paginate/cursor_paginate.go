package paginate

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/tcampbppu/server/app/models"
	"github.com/tcampbppu/server/app/traits/harmony"
	"github.com/tcampbppu/server/database"
	"github.com/tcampbppu/server/pkg/helpers"
	"gorm.io/gorm"

	"encoding/base64"
	"encoding/json"
	"time"
)

type Cursor map[string]interface{}

func CursorPaginate(c *fiber.Ctx, db *gorm.DB) error {
	products := []models.Product{}

	perPage := c.Query("per_page", "10")
	sortOrder := c.Query("sort_order", "desc")
	cursor := c.Query("cursor", "")

	limit, err := strconv.ParseInt(perPage, 10, 64)
	if limit < 1 || limit > 100 {
		limit = 10
	}
	if err != nil {
		return c.Status(500).JSON("Invalid per_page option")
	}

	isFirstPage := cursor == ""
	pointsNext := false

	query := database.DB
	if cursor != "" {
		decodedCursor, err := decodeCursor(cursor)
		if err != nil {
			fmt.Println(err)
			return c.SendStatus(500)
		}
		pointsNext = decodedCursor["points_next"] == true

		operator, order := getPaginationOperator(pointsNext, sortOrder)
		whereStr := fmt.Sprintf("(created_at %s ? OR (created_at = ? AND id %s ?))", operator, operator)
		query = query.Where(whereStr, decodedCursor["created_at"], decodedCursor["created_at"], decodedCursor["id"])
		if order != "" {
			sortOrder = order
		}
	}
	query.Order("created_at " + sortOrder).Limit(int(limit) + 1).Find(&products)
	hasPagination := len(products) > int(limit)

	if hasPagination {
		products = products[:limit]
	}

	if !isFirstPage && !pointsNext {
		products = helpers.Reverse(products)
	}

	pageInfo := calculatePagination(isFirstPage, hasPagination, int(limit), products, pointsNext)

	response := harmony.ResponseDTO{
		Type:    "success",
		Success: true,
		Data:    products,
		Cursor:  pageInfo,
	}
	return harmony.Harmony(c, fiber.StatusOK, response)
}

func calculatePagination(isFirstPage bool, hasPagination bool, limit int, products []models.Product, pointsNext bool) harmony.CursorPaginationInfo {
	pagination := harmony.CursorPaginationInfo{}
	nextCur := Cursor{}
	prevCur := Cursor{}
	if isFirstPage {
		if hasPagination {
			nextCur := CreateCursor(products[limit-1].ID, products[limit-1].CreatedAt, true)
			pagination = GeneratePager(nextCur, nil)
		}
	} else {
		if pointsNext {
			// if pointing next, it always has prev but it might not have next
			if hasPagination {
				nextCur = CreateCursor(products[limit-1].ID, products[limit-1].CreatedAt, true)
			}
			prevCur = CreateCursor(products[0].ID, products[0].CreatedAt, false)
			pagination = GeneratePager(nextCur, prevCur)
		} else {
			// this is case of prev, there will always be nest, but prev needs to be calculated
			nextCur = CreateCursor(products[limit-1].ID, products[limit-1].CreatedAt, true)
			if hasPagination {
				prevCur = CreateCursor(products[0].ID, products[0].CreatedAt, false)
			}
			pagination = GeneratePager(nextCur, prevCur)
		}
	}
	return pagination
}

func getPaginationOperator(pointsNext bool, sortOrder string) (string, string) {
	if pointsNext && sortOrder == "asc" {
		return ">", ""
	}
	if pointsNext && sortOrder == "desc" {
		return "<", ""
	}
	if !pointsNext && sortOrder == "asc" {
		return "<", "desc"
	}
	if !pointsNext && sortOrder == "desc" {
		return ">", "asc"
	}

	return "", ""
}

func encodeCursor(cursor Cursor) string {
	if len(cursor) == 0 {
		return ""
	}
	serializedCursor, err := json.Marshal(cursor)
	if err != nil {
		return ""
	}
	encodedCursor := base64.StdEncoding.EncodeToString(serializedCursor)
	return encodedCursor
}

func decodeCursor(cursor string) (Cursor, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return nil, err
	}

	var cur Cursor
	if err := json.Unmarshal(decodedCursor, &cur); err != nil {
		return nil, err
	}
	return cur, nil
}

func CreateCursor(id uint, createdAt time.Time, pointsNext bool) Cursor {
	return Cursor{
		"id":          id,
		"created_at":  createdAt,
		"points_next": pointsNext,
	}
}

func GeneratePager(next Cursor, prev Cursor) harmony.CursorPaginationInfo {
	return harmony.CursorPaginationInfo{
		NextCursor: encodeCursor(next),
		PrevCursor: encodeCursor(prev),
	}
}
