package tables

import (
	supa "github.com/nedpals/supabase-go"
)

type SupabaseStore struct {
	client *supa.Client
}

func NewSupabaseStore(url string, key string) *SupabaseStore {
	client := supa.CreateClient(url, key)

	return &SupabaseStore{
		client,
	}
}
