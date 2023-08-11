package service

import (
	"app/model"
	"app/repository"
	"app/transaction"
	"errors"
	"log"
)

type IService interface {
	Get(req model.GetRequest) ([]model.DataResponse, error)
	Add(req model.AddRequest) (error, *int64)
	Update(req model.UpdateRequest) error
	Delete(req model.DeleteRequest) error
}

type Service struct {
	r repository.IRepository
	t transaction.ITransaction
}

func NewService(r repository.IRepository, t transaction.ITransaction) IService {
	return &Service{r: r, t: t}
}

func (s *Service) Get(req model.GetRequest) ([]model.DataResponse, error) {
	data, err := s.r.Get(req)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	//
	// put your business logic
	//
	response := []model.DataResponse{}
	for _, i := range data {
		var res model.DataResponse
		//
		// put your business logic
		//
		res.ID = i.ID
		res.Name = i.Name
		res.Type = i.Type
		res.Detail = i.Detail
		res.Url = i.Url
		response = append(response, res)
	}

	return response, nil
}

func (s *Service) Add(req model.AddRequest) (error, *int64) {
	//
	// put your business logic
	//
	err, lastID := s.r.Add(req)
	if err != nil {
		log.Println(err.Error())
		return err, lastID
	}

	err = s.t.Log(
		req.Name, req.Type,
		req.Detail,
		req.Url,
		req.User,
		req.Idrm,
		"ADD")
	if err != nil {
		log.Println(err.Error())
		return err, lastID
	}

	return nil, lastID
}

func (s *Service) Update(req model.UpdateRequest) error {
	// ตรวจสอบว่าค่าใน JSON ถูกระบุหรือไม่
	if req.Name == "" && req.Type == "" && req.Detail == "" && req.Url == "" {
		return errors.New("no fields to update")
	}

	// สร้าง dynamic SQL query
	query := "UPDATE beer SET "
	var args []interface{}

	if req.Name != "" {
		query += "name = ?, "
		args = append(args, req.Name)
	}

	if req.Type != "" {
		query += "type = ?, "
		args = append(args, req.Type)
	}

	if req.Detail != "" {
		query += "detail = ?, "
		args = append(args, req.Detail)
	}

	if req.Url != "" {
		query += "url = ?, "
		args = append(args, req.Url)
	}
	query = query[:len(query)-2] + " WHERE id = ?"
	args = append(args, req.ID)

	// ส่งค่าที่ต้องการอัปเดตไปยัง Repository
	err := s.r.Update(query, args...)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.t.Log(req.Name, req.Type, req.Detail, req.Url, req.User, req.Idrm, "Update")
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (s *Service) Delete(req model.DeleteRequest) error {

	err := s.r.Delete(req)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	err = s.t.Log(req.Name, req.Type,
		req.Detail,
		req.Url,
		req.User,
		req.Idrm,
		"Delete")
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
