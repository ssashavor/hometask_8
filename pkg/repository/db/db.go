package repository

import (
	"database/sql"
	"github.com/ssashavor/homework_8/pkg/model"
)

type ContactsRepositoryInDB struct {
	db *sql.DB
}

func NewContactRepositoryInDB(db *sql.DB) *ContactsRepositoryInDB {
	return &ContactsRepositoryInDB{
		db: db,
	}
}

func (r *ContactsRepositoryInDB) Save(contact model.Contact) (model.Contact, error) {
	err := r.db.QueryRow("INSERT INTO contact(firstname,lastname,phone,email) VALUES($1,$2,$3,$4) returning id", contact.FirstName, contact.LastName, contact.Phone, contact.Email).Scan(&contact.ID)

	if err != nil {
		return model.Contact{}, err
	}
	return contact, nil
}

func (r *ContactsRepositoryInDB) ListAll() (contact []model.Contact, err error) {
	rows, err := r.db.Query("select id, firstname,lastname,phone,email from contact")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		p := model.Contact{}
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone, &p.Email)
		if err != nil {
			return nil, err
		}
		contact = append(contact, p)
	}
	return contact, nil
}

func (r *ContactsRepositoryInDB) GetByID(id uint) (p model.Contact, err error) {
	row := r.db.QueryRow("select id, firstname,lastname,phone,email from contact where id = $1", id)
	err = row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone, &p.Email)
	if err != nil {
		return model.Contact{}, err
	}
	return p, nil
}

func (r *ContactsRepositoryInDB) GetByPhone(phone string) (p model.Contact, err error) {
	row := r.db.QueryRow("select id, firstname,lastname,phone,email from contact where phone = $1", phone)
	err = row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone, &p.Email)
	if err != nil {
		return model.Contact{}, err
	}
	return p, nil
}

func (r *ContactsRepositoryInDB) GetByEmail(email string) (p model.Contact, err error) {
	row := r.db.QueryRow("select id, firstname,lastname,phone,email from contact where email = $1", email)
	err = row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone, &p.Email)
	if err != nil {
		return model.Contact{}, err
	}
	return p, nil
}

func (r *ContactsRepositoryInDB) SearchByName(n string) (contact []model.Contact, err error) {
	row, err := r.db.Query("select id, firstname,lastname,phone,email from contact where firstname = $1", n)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		p := model.Contact{}
		err := row.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Phone, &p.Email)
		if err != nil {
			return nil, err
		}
		contact = append(contact, p)
	}
	return contact, nil
}

func (r *ContactsRepositoryInDB) Delete(id uint) error {
	_, err := r.db.Exec("delete from contact where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
