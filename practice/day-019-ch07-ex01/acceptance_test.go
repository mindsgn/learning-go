package main

import (
	"reflect"
	"testing"
)

func TestTeamAndLeagueTypes(t *testing.T) {
	team := reflect.TypeOf(Team{})
	if team.NumField() != 2 {
		t.Fatalf("Team should have 2 fields, got %d", team.NumField())
	}
	if field := team.Field(0); field.Name != "Name" || field.Type.Kind() != reflect.String {
		t.Fatalf("Team.Name should be a string, got %s %s", field.Name, field.Type)
	}
	if field := team.Field(1); field.Name != "Players" || field.Type.Kind() != reflect.Slice || field.Type.Elem().Kind() != reflect.String {
		t.Fatalf("Team.Players should be []string, got %s %s", field.Name, field.Type)
	}
	league := reflect.TypeOf(League{})
	if league.NumField() != 3 {
		t.Fatalf("League should have 3 fields, got %d", league.NumField())
	}
	checks := map[string]reflect.Kind{"Name": reflect.String, "Teams": reflect.Map, "Wins": reflect.Map}
	for i := 0; i < league.NumField(); i++ {
		field := league.Field(i)
		kind, ok := checks[field.Name]
		if !ok {
			t.Fatalf("unexpected field %s on League", field.Name)
		}
		if field.Type.Kind() != kind {
			t.Fatalf("League.%s should be %v, got %v", field.Name, kind, field.Type.Kind())
		}
	}
}
