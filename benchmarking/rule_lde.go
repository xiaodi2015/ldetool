
/*
 This file was autogenerated via
 --------------------------------------------------------
 ldetool generate --little-endian --package main rule.lde
 --------------------------------------------------------
 do not touch it with bare hands!
*/

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"unsafe"
)

var avatarsUscoreLastcheckLsbrck = []byte("avatars_lastcheck[")
var cavatarUscoreLastmodLsbrck = []byte("cavatar_lastmod[")
var creatorLsbrck = []byte("creator[")
var fLAGSLsbrckSetColon = []byte("FLAGS[set:")
var from = []byte("from")
var reqidSpace = []byte("reqid '")
var spaceActivityEq = []byte(" Activity=")
var spaceAnkUscoreVerLsbrck = []byte(" ank_ver[")
var spaceCreatedLsbrck = []byte(" created[")
var spaceFIELDSLsbrckChangedColon = []byte(" FIELDS[changed:")
var spaceGeoEqLbraceLatColonSpace = []byte(" Geo={Lat: ")
var spaceListUscoreVerLsbrck = []byte(" list_ver[")
var spaceLocationLsbrck = []byte(" location[")
var spacePRESENCESpaceUidEq = []byte(" PRESENCE uid=")
var spaceRegionsLsbrck = []byte(" regions[")
var violationLsbrck = []byte("violation[")

// Line ...
type Line struct {
	rest  []byte
	Name  []byte
	Count []byte
}

// Extract ...
func (p *Line) Extract(line []byte) (bool, error) {
	p.rest = line
	var pos int

	// Take until '|' as Name(string)
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Name = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Looking for '|' and then pass it
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until '|' as Count(string)
	pos = -1
	for i, char := range p.rest {
		if char == '|' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Count = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	return true, nil
}

// CRMod ...
type CRMod struct {
	rest          []byte
	Time          []byte
	ChatID        uint64
	ReqID         []byte
	UIN           []byte
	FlagsSet      []byte
	FlagsUnset    []byte
	FieldsChanged []byte
	AnkVer        struct {
		Valid bool
		Value []byte
	}
	ListVer struct {
		Valid bool
		Value []byte
	}
	Name  []byte
	About struct {
		Valid bool
		Value []byte
	}
	Rules struct {
		Valid bool
		Value []byte
	}
	Nick struct {
		Valid bool
		Value []byte
	}
	Location struct {
		Valid bool
		Value []byte
	}
	Stamp struct {
		Valid bool
		Value []byte
	}
	Regions struct {
		Valid bool
		Value []byte
	}
	Flags struct {
		Valid bool
		Value []byte
	}
	Created int64
	Creator struct {
		Valid bool
		Value []byte
	}
	AvatarLastCheck struct {
		Valid bool
		Value int64
	}
	AvatarsLastMod struct {
		Valid bool
		Value int64
	}
	Origin    []byte
	Drugs     int16
	Spam      int16
	Pron      int16
	Violation struct {
		Valid bool
		Value int16
	}
	AbuseOther struct {
		Valid bool
		Value int16
	}
}

// Extract ...
func (p *CRMod) Extract(line []byte) (bool, error) {
	p.rest = line
	var err error
	var pos int
	var rest1 []byte
	var tmp []byte
	var tmpInt int64
	var tmpUint uint64

	// Checks if the rest starts with '[' and pass it
	if len(p.rest) >= 1 && p.rest[0] == '[' {
		p.rest = p.rest[1:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m)` is expected to start with \033[1m%s\033[0m", string(p.rest), '[')
	}

	// Looking for ' ' and then pass it
	pos = bytes.IndexByte(p.rest, ' ')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`", ' ', string(p.rest))
	}

	// Take until ']' as Time(string)
	pos = bytes.IndexByte(p.rest, ']')
	if pos >= 0 {
		p.Time = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Time", ']', string(p.rest))
	}

	// Looking for '[' and then pass it
	pos = bytes.IndexByte(p.rest, '[')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`", '[', string(p.rest))
	}

	// Take until '.' as ChatID(uint64)
	pos = bytes.IndexByte(p.rest, '.')
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field ChatID", '.', string(p.rest))
	}
	if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&tmp)), 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.ChatID = uint64(tmpUint)

	// Looking for "reqid '" and then pass it
	pos = bytes.Index(p.rest, reqidSpace)
	if pos >= 0 {
		p.rest = p.rest[pos+len(reqidSpace):]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`", reqidSpace, string(p.rest))
	}

	// Take until '\'' as ReqID(string)
	pos = bytes.IndexByte(p.rest, '\'')
	if pos >= 0 {
		p.ReqID = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field ReqID", '\'', string(p.rest))
	}

	// Looking for "from" and then pass it
	pos = bytes.Index(p.rest, from)
	if pos >= 0 {
		p.rest = p.rest[pos+len(from):]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`", from, string(p.rest))
	}

	// Looking for '(' and then pass it
	pos = bytes.IndexByte(p.rest, '(')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find \033[1m%c\033[0m in `\033[1m%s\033[0m`", '(', string(p.rest))
	}

	// Take until ')' as UIN(string)
	pos = bytes.IndexByte(p.rest, ')')
	if pos >= 0 {
		p.UIN = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field UIN", ')', string(p.rest))
	}

	// Looking for "FLAGS[set:" and then pass it
	pos = bytes.Index(p.rest, fLAGSLsbrckSetColon)
	if pos >= 0 {
		p.rest = p.rest[pos+len(fLAGSLsbrckSetColon):]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%s\033[0m` in `\033[1m%s\033[0m`", fLAGSLsbrckSetColon, string(p.rest))
	}

	// Take until ',' as FlagsSet(string)
	pos = bytes.IndexByte(p.rest, ',')
	if pos >= 0 {
		p.FlagsSet = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field FlagsSet", ',', string(p.rest))
	}

	// Checks if the rest starts with `" unset:"` and pass it
	if len(p.rest) >= 7 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffffffffff == 0x3a7465736e7520 {
		p.rest = p.rest[7:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " unset:")
	}

	// Take until ']' as FlagsUnset(string)
	pos = bytes.IndexByte(p.rest, ']')
	if pos >= 0 {
		p.FlagsUnset = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field FlagsUnset", ']', string(p.rest))
	}

	// Checks if the rest starts with `" FIELDS[changed:"` and pass it
	if bytes.HasPrefix(p.rest, spaceFIELDSLsbrckChangedColon) {
		p.rest = p.rest[len(spaceFIELDSLsbrckChangedColon):]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " FIELDS[changed:")
	}

	// Take until ']' as FieldsChanged(string)
	pos = bytes.IndexByte(p.rest, ']')
	if pos >= 0 {
		p.FieldsChanged = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field FieldsChanged", ']', string(p.rest))
	}
	rest1 = p.rest

	// Checks if the rest starts with `" ank_ver["` and pass it
	if bytes.HasPrefix(rest1, spaceAnkUscoreVerLsbrck) {
		rest1 = rest1[len(spaceAnkUscoreVerLsbrck):]
	} else {
		p.AnkVer.Valid = false
		goto crmodAnkVerLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.AnkVer.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.AnkVer.Valid = false
		goto crmodAnkVerLabel
	}
	p.AnkVer.Valid = true
	p.rest = rest1
crmodAnkVerLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" list_ver["` and pass it
	if bytes.HasPrefix(rest1, spaceListUscoreVerLsbrck) {
		rest1 = rest1[len(spaceListUscoreVerLsbrck):]
	} else {
		p.ListVer.Valid = false
		goto crmodListVerLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.ListVer.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.ListVer.Valid = false
		goto crmodListVerLabel
	}
	p.ListVer.Valid = true
	p.rest = rest1
crmodListVerLabel:

	// Checks if the rest starts with `" name["` and pass it
	if len(p.rest) >= 6 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffffffff == 0x5b656d616e20 {
		p.rest = p.rest[6:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " name[")
	}

	// Take until ']' as Name(string)
	pos = bytes.IndexByte(p.rest, ']')
	if pos >= 0 {
		p.Name = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Name", ']', string(p.rest))
	}
	rest1 = p.rest

	// Checks if the rest starts with `" about["` and pass it
	if len(rest1) >= 7 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffffff == 0x5b74756f626120 {
		rest1 = rest1[7:]
	} else {
		p.About.Valid = false
		goto crmodAboutLabel
	}

	// Take until ']' as Value(string)
	pos = -1
	for i, char := range rest1 {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.About.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.About.Valid = false
		goto crmodAboutLabel
	}
	p.About.Valid = true
	p.rest = rest1
crmodAboutLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" rules["` and pass it
	if len(rest1) >= 7 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffffff == 0x5b73656c757220 {
		rest1 = rest1[7:]
	} else {
		p.Rules.Valid = false
		goto crmodRulesLabel
	}

	// Take until ']' as Value(string)
	pos = -1
	for i, char := range rest1 {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Rules.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Rules.Valid = false
		goto crmodRulesLabel
	}
	p.Rules.Valid = true
	p.rest = rest1
crmodRulesLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" nick["` and pass it
	if len(rest1) >= 6 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffff == 0x5b6b63696e20 {
		rest1 = rest1[6:]
	} else {
		p.Nick.Valid = false
		goto crmodNickLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.Nick.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Nick.Valid = false
		goto crmodNickLabel
	}
	p.Nick.Valid = true
	p.rest = rest1
crmodNickLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" location["` and pass it
	if bytes.HasPrefix(rest1, spaceLocationLsbrck) {
		rest1 = rest1[len(spaceLocationLsbrck):]
	} else {
		p.Location.Valid = false
		goto crmodLocationLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.Location.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Location.Valid = false
		goto crmodLocationLabel
	}
	p.Location.Valid = true
	p.rest = rest1
crmodLocationLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" stamp["` and pass it
	if len(rest1) >= 7 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffffff == 0x5b706d61747320 {
		rest1 = rest1[7:]
	} else {
		p.Stamp.Valid = false
		goto crmodStampLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.Stamp.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Stamp.Valid = false
		goto crmodStampLabel
	}
	p.Stamp.Valid = true
	p.rest = rest1
crmodStampLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" regions["` and pass it
	if bytes.HasPrefix(rest1, spaceRegionsLsbrck) {
		rest1 = rest1[len(spaceRegionsLsbrck):]
	} else {
		p.Regions.Valid = false
		goto crmodRegionsLabel
	}

	// Take until ']' as Value(string)
	pos = -1
	for i, char := range rest1 {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Regions.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Regions.Valid = false
		goto crmodRegionsLabel
	}
	p.Regions.Valid = true
	p.rest = rest1
crmodRegionsLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" flags["` and pass it
	if len(rest1) >= 7 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffffff == 0x5b7367616c6620 {
		rest1 = rest1[7:]
	} else {
		p.Flags.Valid = false
		goto crmodFlagsLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.Flags.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Flags.Valid = false
		goto crmodFlagsLabel
	}
	p.Flags.Valid = true
	p.rest = rest1
crmodFlagsLabel:

	// Checks if the rest starts with `" created["` and pass it
	if bytes.HasPrefix(p.rest, spaceCreatedLsbrck) {
		p.rest = p.rest[len(spaceCreatedLsbrck):]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " created[")
	}

	// Take until '=' as Created(int64)
	pos = bytes.IndexByte(p.rest, '=')
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Created", '=', string(p.rest))
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Created = int64(tmpInt)
	rest1 = p.rest

	// Looking for "creator[" and then pass it
	pos = bytes.Index(rest1, creatorLsbrck)
	if pos >= 0 {
		rest1 = rest1[pos+len(creatorLsbrck):]
	} else {
		p.Creator.Valid = false
		goto crmodCreatorLabel
	}

	// Take until ']' as Value(string)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		p.Creator.Value = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Creator.Valid = false
		goto crmodCreatorLabel
	}
	p.Creator.Valid = true
	p.rest = rest1
crmodCreatorLabel:
	rest1 = p.rest

	// Looking for "avatars_lastcheck[" and then pass it
	pos = bytes.Index(rest1, avatarsUscoreLastcheckLsbrck)
	if pos >= 0 {
		rest1 = rest1[pos+len(avatarsUscoreLastcheckLsbrck):]
	} else {
		p.AvatarLastCheck.Valid = false
		goto crmodAvatarLastCheckLabel
	}

	// Take until ']' as Value(int64)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.AvatarLastCheck.Valid = false
		goto crmodAvatarLastCheckLabel
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.AvatarLastCheck.Value = int64(tmpInt)
	p.AvatarLastCheck.Valid = true
	p.rest = rest1
crmodAvatarLastCheckLabel:
	rest1 = p.rest

	// Looking for "cavatar_lastmod[" and then pass it
	pos = bytes.Index(rest1, cavatarUscoreLastmodLsbrck)
	if pos >= 0 {
		rest1 = rest1[pos+len(cavatarUscoreLastmodLsbrck):]
	} else {
		p.AvatarsLastMod.Valid = false
		goto crmodAvatarsLastModLabel
	}

	// Take until ']' as Value(int64)
	pos = bytes.IndexByte(rest1, ']')
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.AvatarsLastMod.Valid = false
		goto crmodAvatarsLastModLabel
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.AvatarsLastMod.Value = int64(tmpInt)
	p.AvatarsLastMod.Valid = true
	p.rest = rest1
crmodAvatarsLastModLabel:

	// Checks if the rest starts with `" origin["` and pass it
	if len(p.rest) >= 8 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffffffffffff == 0x5b6e696769726f20 {
		p.rest = p.rest[8:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " origin[")
	}

	// Take until ']' as Origin(string)
	pos = -1
	for i, char := range p.rest {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		p.Origin = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Origin", ']', string(p.rest))
	}

	// Checks if the rest starts with `" abuse"` and pass it
	if len(p.rest) >= 6 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffffffff == 0x657375626120 {
		p.rest = p.rest[6:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " abuse")
	}

	// Checks if rest[1:] starts with `"drugs["` and pass it
	if len(p.rest)-1 >= 6 && *(*uint64)(unsafe.Pointer(&p.rest[1]))&0xffffffffffff == 0x5b7367757264 {
		p.rest = p.rest[7:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest[1:]), "drugs[")
	}

	// Take until ']' as Drugs(int16)
	pos = -1
	for i, char := range p.rest {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Drugs", ']', string(p.rest))
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 16); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Drugs = int16(tmpInt)

	// Checks if the rest starts with `" abuse"` and pass it
	if len(p.rest) >= 6 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffffffff == 0x657375626120 {
		p.rest = p.rest[6:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " abuse")
	}

	// Checks if rest[1:] starts with `"spam["` and pass it
	if len(p.rest)-1 >= 5 && *(*uint64)(unsafe.Pointer(&p.rest[1]))&0xffffffffff == 0x5b6d617073 {
		p.rest = p.rest[6:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest[1:]), "spam[")
	}

	// Take until ']' as Spam(int16)
	pos = -1
	for i, char := range p.rest {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Spam", ']', string(p.rest))
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 16); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Spam = int16(tmpInt)

	// Checks if the rest starts with `" abuse"` and pass it
	if len(p.rest) >= 6 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffffffff == 0x657375626120 {
		p.rest = p.rest[6:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " abuse")
	}

	// Checks if rest[1:] starts with `"porno["` and pass it
	if len(p.rest)-1 >= 6 && *(*uint64)(unsafe.Pointer(&p.rest[1]))&0xffffffffffff == 0x5b6f6e726f70 {
		p.rest = p.rest[7:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest[1:]), "porno[")
	}

	// Take until ']' as Pron(int16)
	pos = -1
	for i, char := range p.rest {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		tmp = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field Pron", ']', string(p.rest))
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 16); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Pron = int16(tmpInt)
	rest1 = p.rest

	// Checks if the rest starts with `" abuse"` and pass it
	if len(rest1) >= 6 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffff == 0x657375626120 {
		rest1 = rest1[6:]
	} else {
		p.Violation.Valid = false
		goto crmodViolationLabel
	}

	// Checks if rest[1:] starts with `"violation["` and pass it
	if len(rest1[1:]) >= 1+len(violationLsbrck) && bytes.HasPrefix(rest1[1:], violationLsbrck) {
		rest1 = rest1[len(violationLsbrck)+1:]
	} else {
		p.Violation.Valid = false
		goto crmodViolationLabel
	}

	// Take until ']' as Value(int16)
	pos = -1
	for i, char := range rest1 {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Violation.Valid = false
		goto crmodViolationLabel
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 16); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Violation.Value = int16(tmpInt)
	p.Violation.Valid = true
	p.rest = rest1
crmodViolationLabel:
	rest1 = p.rest

	// Checks if the rest starts with `" abuse"` and pass it
	if len(rest1) >= 6 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffff == 0x657375626120 {
		rest1 = rest1[6:]
	} else {
		p.AbuseOther.Valid = false
		goto crmodAbuseOtherLabel
	}

	// Checks if rest[1:] starts with `"other["` and pass it
	if len(rest1)-1 >= 6 && *(*uint64)(unsafe.Pointer(&rest1[1]))&0xffffffffffff == 0x5b726568746f {
		rest1 = rest1[7:]
	} else {
		p.AbuseOther.Valid = false
		goto crmodAbuseOtherLabel
	}

	// Take until ']' as Value(int16)
	pos = -1
	for i, char := range rest1 {
		if char == ']' {
			pos = i
			break
		}
	}
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.AbuseOther.Valid = false
		goto crmodAbuseOtherLabel
	}
	if tmpInt, err = strconv.ParseInt(*(*string)(unsafe.Pointer(&tmp)), 10, 16); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.AbuseOther.Value = int16(tmpInt)
	p.AbuseOther.Valid = true
	p.rest = rest1
crmodAbuseOtherLabel:

	return true, nil
}

// GetAnkVerValue ...
func (p *CRMod) GetAnkVerValue() (res []byte) {
	if p.AnkVer.Valid {
		res = p.AnkVer.Value
	}
	return
}

// GetListVerValue ...
func (p *CRMod) GetListVerValue() (res []byte) {
	if p.ListVer.Valid {
		res = p.ListVer.Value
	}
	return
}

// GetAboutValue ...
func (p *CRMod) GetAboutValue() (res []byte) {
	if p.About.Valid {
		res = p.About.Value
	}
	return
}

// GetRulesValue ...
func (p *CRMod) GetRulesValue() (res []byte) {
	if p.Rules.Valid {
		res = p.Rules.Value
	}
	return
}

// GetNickValue ...
func (p *CRMod) GetNickValue() (res []byte) {
	if p.Nick.Valid {
		res = p.Nick.Value
	}
	return
}

// GetLocationValue ...
func (p *CRMod) GetLocationValue() (res []byte) {
	if p.Location.Valid {
		res = p.Location.Value
	}
	return
}

// GetStampValue ...
func (p *CRMod) GetStampValue() (res []byte) {
	if p.Stamp.Valid {
		res = p.Stamp.Value
	}
	return
}

// GetRegionsValue ...
func (p *CRMod) GetRegionsValue() (res []byte) {
	if p.Regions.Valid {
		res = p.Regions.Value
	}
	return
}

// GetFlagsValue ...
func (p *CRMod) GetFlagsValue() (res []byte) {
	if p.Flags.Valid {
		res = p.Flags.Value
	}
	return
}

// GetCreatorValue ...
func (p *CRMod) GetCreatorValue() (res []byte) {
	if p.Creator.Valid {
		res = p.Creator.Value
	}
	return
}

// GetAvatarLastCheckValue ...
func (p *CRMod) GetAvatarLastCheckValue() (res int64) {
	if p.AvatarLastCheck.Valid {
		res = p.AvatarLastCheck.Value
	}
	return
}

// GetAvatarsLastModValue ...
func (p *CRMod) GetAvatarsLastModValue() (res int64) {
	if p.AvatarsLastMod.Valid {
		res = p.AvatarsLastMod.Value
	}
	return
}

// GetViolationValue ...
func (p *CRMod) GetViolationValue() (res int16) {
	if p.Violation.Valid {
		res = p.Violation.Value
	}
	return
}

// GetAbuseOtherValue ...
func (p *CRMod) GetAbuseOtherValue() (res int16) {
	if p.AbuseOther.Valid {
		res = p.AbuseOther.Value
	}
	return
}

// Presence ...
type Presence struct {
	rest []byte
	Time []byte
	UID  []byte
	UA   []byte
	Geo  struct {
		Valid bool
		Lat   float64
		Lon   float64
	}
	Activity uint8
}

// Extract ...
func (p *Presence) Extract(line []byte) (bool, error) {
	p.rest = line
	var err error
	var pos int
	var rest1 []byte
	var tmp []byte
	var tmpFloat float64
	var tmpUint uint64

	// Looking for ' ' and then pass it
	pos = bytes.IndexByte(p.rest, ' ')
	if pos >= 0 {
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Take until ']' as Time(string)
	pos = bytes.IndexByte(p.rest, ']')
	if pos >= 0 {
		p.Time = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, nil
	}

	// Checks if the rest starts with `" PRESENCE uid="` and pass it
	if bytes.HasPrefix(p.rest, spacePRESENCESpaceUidEq) {
		p.rest = p.rest[len(spacePRESENCESpaceUidEq):]
	} else {
		return false, nil
	}

	// Take until ' ' as UID(string)
	pos = bytes.IndexByte(p.rest, ' ')
	if pos >= 0 {
		p.UID = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field UID", ' ', string(p.rest))
	}

	// Checks if the rest starts with `"ua='"` and pass it
	if len(p.rest) >= 4 && *(*uint64)(unsafe.Pointer(&p.rest[0]))&0xffffffff == 0x273d6175 {
		p.rest = p.rest[4:]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), "ua='")
	}

	// Take until '\'' as UA(string)
	pos = bytes.IndexByte(p.rest, '\'')
	if pos >= 0 {
		p.UA = p.rest[:pos]
		p.rest = p.rest[pos+1:]
	} else {
		return false, fmt.Errorf("Cannot find `\033[1m%c\033[0m` in `\033[1m%s\033[0m` to bound data for field UA", '\'', string(p.rest))
	}
	rest1 = p.rest

	// Checks if the rest starts with `" Geo={Lat: "` and pass it
	if bytes.HasPrefix(rest1, spaceGeoEqLbraceLatColonSpace) {
		rest1 = rest1[len(spaceGeoEqLbraceLatColonSpace):]
	} else {
		p.Geo.Valid = false
		goto presenceGeoLabel
	}

	// Take until ',' as Lat(float64)
	pos = bytes.IndexByte(rest1, ',')
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Geo.Valid = false
		goto presenceGeoLabel
	}
	if tmpFloat, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&tmp)), 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Geo.Lat = float64(tmpFloat)

	// Checks if the rest starts with `" Lon: "` and pass it
	if len(rest1) >= 6 && *(*uint64)(unsafe.Pointer(&rest1[0]))&0xffffffffffff == 0x203a6e6f4c20 {
		rest1 = rest1[6:]
	} else {
		p.Geo.Valid = false
		goto presenceGeoLabel
	}

	// Take until '}' as Lon(float64)
	pos = bytes.IndexByte(rest1, '}')
	if pos >= 0 {
		tmp = rest1[:pos]
		rest1 = rest1[pos+1:]
	} else {
		p.Geo.Valid = false
		goto presenceGeoLabel
	}
	if tmpFloat, err = strconv.ParseFloat(*(*string)(unsafe.Pointer(&tmp)), 64); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(tmp), err)
	}
	p.Geo.Lon = float64(tmpFloat)
	p.Geo.Valid = true
	p.rest = rest1
presenceGeoLabel:

	// Checks if the rest starts with `" Activity="` and pass it
	if bytes.HasPrefix(p.rest, spaceActivityEq) {
		p.rest = p.rest[len(spaceActivityEq):]
	} else {
		return false, fmt.Errorf("`\033[1m%s\033[0m` is expected to start with `\033[1m%s\033[0m`", string(p.rest), " Activity=")
	}

	// Take the rest as Activity(uint8)
	if tmpUint, err = strconv.ParseUint(*(*string)(unsafe.Pointer(&p.rest)), 10, 8); err != nil {
		return false, fmt.Errorf("Cannot parse `%s`: %s", string(p.rest), err)
	}
	p.Activity = uint8(tmpUint)
	p.rest = p.rest[len(p.rest):]
	return true, nil
}

// GetGeoLat ...
func (p *Presence) GetGeoLat() (res float64) {
	if p.Geo.Valid {
		res = p.Geo.Lat
	}
	return
}

// GetGeoLon ...
func (p *Presence) GetGeoLon() (res float64) {
	if p.Geo.Valid {
		res = p.Geo.Lon
	}
	return
}
