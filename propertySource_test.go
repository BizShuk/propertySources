package main

import (
	pSource "propertysources/propertysource"
	"reflect"
	"testing"
)

func TestPropertiesSources_load(t *testing.T) {
	type fields struct {
		plist []pSource.Properties
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PropertiesSources{
				plist: tt.fields.plist,
			}
			p.load()
		})
	}
}

func TestPropertiesSources_appendProperties(t *testing.T) {
	type fields struct {
		plist []pSource.Properties
	}
	type args struct {
		properties pSource.Properties
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PropertiesSources{
				plist: tt.fields.plist,
			}
			p.appendProperties(tt.args.properties)
		})
	}
}

func TestPropertiesSources_Get(t *testing.T) {
	type fields struct {
		plist []pSource.Properties
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantVal string
		wantOk  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PropertiesSources{
				plist: tt.fields.plist,
			}
			gotVal, gotOk := p.Get(tt.args.key)
			if gotVal != tt.wantVal {
				t.Errorf("PropertiesSources.Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("PropertiesSources.Get() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *PropertiesSources
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
