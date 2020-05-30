package chatbotdb

import (
	"strconv"
	"unicode"

	chatbotbase "github.com/zhs007/chatbot/base"
	chatbotpb "github.com/zhs007/chatbot/chatbotpb"
)

// AppServDBKeyPrefix - This is the prefix of AppServDBKey
const AppServDBKeyPrefix = "as:"

// makeAppServKey - Generate a database key via token
func makeAppServKey(token string) string {
	return chatbotbase.AppendString(AppServDBKeyPrefix, token)
}

// UserIDDBKey - This is UserIDDBKey
const UserIDDBKey = "uid"

// AppDataDBKey - This is AppDataDBKey
const AppDataDBKey = "appdata"

// UserDBKeyPrefix - This is the prefix of UserDBKey
const UserDBKeyPrefix = "ui:"

// makeUserKey - Generate a database key via uid
func makeUserKey(uid int64) string {
	return chatbotbase.AppendString(UserDBKeyPrefix, strconv.FormatInt(uid, 10))
}

// UserDataDBKeyPrefix - This is the prefix of UserDataDBKey
const UserDataDBKeyPrefix = "ud:"

// makeUserDataKey - Generate a database key via uid
func makeUserDataKey(uid int64) string {
	return chatbotbase.AppendString(UserDataDBKeyPrefix, strconv.FormatInt(uid, 10))
}

// AppUIDDBKeyPrefix - This is the prefix of AppUIDDBKey
const AppUIDDBKeyPrefix = "at:"

// makeAppUID - Generate a database key via apptoken and appuid
func makeAppUID(appToken string, appUID string) string {
	return chatbotbase.AppendString(AppUIDDBKeyPrefix, appToken, ":", appUID)
}

// NoteInfoDBKeyPrefix - This is the prefix of NoteInfo
const NoteInfoDBKeyPrefix = "ni:"

// makeNoteInfoKey - Generate a database key via note name
func makeNoteInfoKey(name string) string {
	return chatbotbase.AppendString(NoteInfoDBKeyPrefix, name)
}

// IsValidNoteName - is valid note name
func IsValidNoteName(name string) bool {
	if name == "" {
		return false
	}

	for _, v := range name {
		if !(unicode.IsLetter(v) || unicode.IsDigit(v)) {
			return false
		}
	}

	return true
}

// NoteNodeDBKeyPrefix - This is the prefix of NoteNode
const NoteNodeDBKeyPrefix = "nn:"

// makeNoteNodeKey - Generate a database key via note name
func makeNoteNodeKey(name string, nodeIndex int64) string {
	return chatbotbase.AppendString(NoteInfoDBKeyPrefix, name,
		":", strconv.FormatInt(nodeIndex, 10))
}

// MergeKeys - return arr0 + arr1
func MergeKeys(arr0 []string, arr1 []string) []string {
	for _, v := range arr1 {
		if chatbotbase.IndexOfArrayString(arr0, v) < 0 {
			arr0 = append(arr0, v)
		}
	}

	return arr0
}

// InsMapKeys - insert into mapKeys
func InsMapKeys(ni *chatbotpb.NoteInfo, keys []string, noteIndex int64) {
	if ni.MapKeys == nil {
		ni.MapKeys = make(map[string]*chatbotpb.NoteKeyInfo)
	}

	for _, v := range keys {
		ks, isok := ni.MapKeys[v]
		if !isok {
			ni.MapKeys[v] = &chatbotpb.NoteKeyInfo{
				Nodes: []int64{noteIndex},
			}
		} else {
			ks.Nodes = append(ks.Nodes, noteIndex)
		}
	}
}

func insKeys(lst []int64, noteIndex int64) []int64 {
	for _, v := range lst {
		if v == noteIndex {
			return lst
		}
	}

	return append(lst, noteIndex)
}

// SearchKeys - search with keys
func SearchKeys(ni *chatbotpb.NoteInfo, keys []string) []int64 {
	var lst []int64

	if ni.MapKeys == nil {
		return nil
	}

	for _, v := range keys {
		ks, isok := ni.MapKeys[v]
		if isok {
			for _, v1 := range ks.Nodes {
				lst = insKeys(lst, v1)
			}
		}
	}

	return lst
}

// RemoveKeys - remove keys
func RemoveKeys(ni *chatbotpb.NoteInfo, keys []string) *chatbotpb.NoteInfo {
	var nkeys []string
	for _, v := range ni.Keys {
		haskey := false
		for _, v1 := range keys {
			if v == v1 {
				haskey = true

				break
			}
		}

		if !haskey {
			nkeys = append(nkeys, v)
		}
	}

	ni.Keys = nkeys

	if ni.MapKeys != nil {
		for _, v := range keys {
			_, isok := ni.MapKeys[v]
			if isok {
				delete(ni.MapKeys, v)
			}
		}
	}

	return ni
}

// RemoveNoteNode - remove notenode
func RemoveNoteNode(ni *chatbotpb.NoteInfo, i int64) *chatbotpb.NoteInfo {
	if ni.MapKeys != nil {
		for _, v := range ni.MapKeys {

			var nn []int64
			for _, v1 := range v.Nodes {
				if v1 != i {
					nn = append(nn, v1)
				}
			}

			v.Nodes = nn
		}
	}

	return ni
}
