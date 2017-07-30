package models
    
import (
    	gorp "gopkg.in/gorp.v1"
)
    
// Likes struct
type Likes struct {
    ID    int64     `db:"id" json:"id"`
    MittyID    int   `db:"mitty_id" json:"mitty_id"`
    EntityType    string `db:"entity_type" json:"entity_type"`
    EntityID    int64     `db:"entity_id" json:"entity_id"`  
}	

// Insert ...
func (s *Likes) Insert(tx gorp.Transaction) error {
	err := tx.Insert(s)
	return err
}

// Update ...
func (s *Likes) Update(tx gorp.Transaction) error {
	_, err := tx.Update(s)
	return err
}

// Delete ...
func (s *Likes) Delete(tx gorp.Transaction) error {
	_, err := tx.Delete(s)
	return err
}

// RemoveLikesByID ...
func RemoveLikesByID(tx *gorp.Transaction, userId int, entityType string, entityID int64)  error  {
	if _, err := tx.Exec("Delete from Likes  WHERE mitty_id = $1 and entity_tyep=$2 and entity_id = $3", userId, entityType, entityID); err != nil {
		return err
	}
	return nil
}