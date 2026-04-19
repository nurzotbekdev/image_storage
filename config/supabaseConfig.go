package config

import (
	"os"

	"github.com/supabase-community/supabase-go"
)

var Supabase *supabase.Client
var SupabaseBucket string

func InitSupabase() {
	client, err := supabase.NewClient(
		os.Getenv("SUPABASE_URL"),
		os.Getenv("SUPABASE_KEY"),
		&supabase.ClientOptions{},
	)

	if err != nil {
		panic("Supabase ulanishida xatolik: " + err.Error())
	}

	Supabase = client
	SupabaseBucket = os.Getenv("SUPABASE_BUCKET")
}
