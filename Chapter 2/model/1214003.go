package dipa

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Listtamu struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Name string `bson:"name" json:"name"`
	Notelp string `bson:"notelp" json:"notelp"`
	Email string `bson:"email" json:"email"`
	Kota string `bson:"kota" json:"kota"`
}

type Undanganrapat struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Lokasi string `bson:"lokasi" json:"lokasi"`
	Agenda string `bson:"agenda" json:"agenda"`
	Peserta string `bson:"peserta" json:"peserta"`
}

type Pesertarapat struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Nama string `bson:"nama" json:"nama"`
	Instansi string `bson:"instansi" json:"instansi"`
	Status string `bson:"status" json:"status"`
}

type Wakturapat struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Hal string `bson:"hal" json:"hal"`
	Materi string `bson:"materi" json:"materi"`
}

type Rapatmulai struct {
	ID         	primitive.ObjectID `bson:"_id,omitempty" json:"_id"`
	Pembicara string `bson:"pembicara" json:"pembicara"`
	Durasi string `bson:"durasi" json:"durasi"`
}