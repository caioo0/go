package mlib

import "testing"

func TestOps(t *testing.T) {
    // TODO: Implement tests for the Manager struct
    mm := NewMusicManager()

    if mm == nil {
        t.Error("newmusicmanager() returned nil")
    }
    if mm.Len() != 0 {
        t.Error("newmusicmanager() returned a non-empty manager")
    }
    m0 := &musicEntry{
        "1","My heart will go on", "Audrey Hepburn", Pop,"http://someurl.com", MP3
    }
    mm.Add(m0)

    if mm.Len() != 1 {
        t.Error("Add() failed")
    }

    m :=mm.Find(m0.Name)
    if m == nil {
        t.Error("Find() failed")
    }
    if m.Id != m0.Id  || m.Name != m0.Name || m.Artist != m0.Artist || m.Genre != m0.Genre || m.Url != m0.Url || m.Type != m0.Type {
         t.Error("Find() failed,found item mismatch.")
    }

    m, err := mm.Get(0)
    if err != nil {
        t.Error("musicmanager.Get() failed",err)
    }

    m = mm.Remove(0)
    if m == nil || mm.Len() != 0 {
        t.Error("musicmanager.Remove() failed",err)
    }
    
}