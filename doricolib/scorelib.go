package doricolib

import (
	"encoding/xml"
	"fmt"
	"io"
)

// Definitions for Dorico's kScoreLibrary
type ScoreLib struct {
	XMLName                                             string              `xml:"kScoreLibrary"`
	FileVersion                                         string              `xml:"fileVersion"`
	Temperaments                                        EntityListContainer `xml:"temperaments"`
	AccidentalSystems                                   EntityListContainer `xml:"accidentalSystems"`
	AccidentalDefinitions                               EntityListContainer `xml:"accidentalDefinitions"`
	TonalitySystemDefinitions                           EntityListContainer `xml:"tonalitySystemDefinitions"`
	Ensembles                                           EntityListContainer `xml:"ensembles"`
	Instruments                                         EntityListContainer `xml:"instruments"`
	InstrumentNames                                     InstrumentNameList  `xml:"instrumentNames"`
	InstrumentFamilies                                  EntityListContainer `xml:"instrumentFamilies"`
	ClefDefinitions                                     EntityListContainer `xml:"clefDefinitions"`
	OttavaLineDefinitions                               EntityListContainer `xml:"ottavaLineDefinitions"`
	NoteheadSetDefinitions                              EntityListContainer `xml:"noteheadSetDefinitions"`
	NoteheadDefinitions                                 EntityListContainer `xml:"noteheadDefinitions"`
	RestSetDefinitions                                  EntityListContainer `xml:"restSetDefinitions"`
	RestDefinitions                                     EntityListContainer `xml:"restDefinitions"`
	FlagSetDefinitions                                  EntityListContainer `xml:"flagSetDefinitions"`
	FlagDefinitions                                     EntityListContainer `xml:"flagDefinitions"`
	GraphicDefinitions                                  EntityListContainer `xml:"graphicDefinitions"`
	BlobDefinitions                                     EntityListContainer `xml:"blobDefinitions"`
	DrawingDefinitions                                  EntityListContainer `xml:"drawingDefinitions"`
	TextDefinitions                                     EntityListContainer `xml:"textDefinitions"`
	GlyphDefinitions                                    EntityListContainer `xml:"glyphDefinitions"`
	CompositeDefinitions                                EntityListContainer `xml:"compositeDefinitions"`
	TempoPresetDefinitions                              EntityListContainer `xml:"tempoPresetDefinitions"`
	RhythmicFeelDefinitions                             EntityListContainer `xml:"rhythmicFeelDefinitions"`
	PenStyles                                           EntityListContainer `xml:"penstyles"`
	FontStyles                                          EntityListContainer `xml:"fontstyles"`
	BrushStyles                                         EntityListContainer `xml:"brushstyles"`
	PagePairDefinitionSets                              EntityListContainer `xml:"pagePairDefinitionSets"`
	PageDimensionPresets                                EntityListContainer `xml:"pageDimensionPresets"`
	ParagraphStyles                                     EntityListContainer `xml:"paragraphStyles"`
	CharacterStyles                                     EntityListContainer `xml:"characterStyles"`
	PlayingTechniques                                   EntityListContainer `xml:"playingTechniques"`
	OrnamentDefinitions                                 EntityListContainer `xml:"ornamentDefinitions"`
	MultiSegmentLineDefinitions                         EntityListContainer `xml:"multiSegmentLineDefinitions"`
	PlayingTechniqueAppearanceCollectionDefinition      EntityListContainer `xml:"playingTechniqueAppearanceCollectionDefinition"`
	ExpressionMaps                                      EntityListContainer `xml:"expressionMapDefinitions"`
	DrumKitNoteMapDefinitions                           EntityListContainer `xml:"drumKitNoteMapDefinitions"`
	RepeatableCompositeCollectionDefinitions            EntityListContainer `xml:"repeatableCompositeCollectionDefinitions"`
	ChordSymbolAppearanceCollectionDefinitions          EntityListContainer `xml:"chordSymbolAppearanceCollectionDefinitions"`
	ChordSymbolAppearanceComponentCollectionDefinitions EntityListContainer `xml:"chordSymbolAppearanceComponentCollectionDefinitions"`
	PercussionKitDefinitionCollectionDefinition         EntityListContainer `xml:"percussionKitDefinitionCollectionDefinition"`
	PercussionInstrumentDataCollectionDefinition        EntityListContainer `xml:"percussionInstrumentDataCollectionDefinition"`
	FrettedInstrumentDataCollectionDefinition           EntityListContainer `xml:"frettedInstrumentDataCollectionDefinition"`
	ChordDiagramCollectionDefinition                    EntityListContainer `xml:"chordDiagramCollectionDefinition"`
	LineAnnotationCollectionDefinition                  EntityListContainer `xml:"lineAnnotationCollectionDefinition"`
	LineBodyStyleCollectionDefinition                   EntityListContainer `xml:"lineBodyStyleCollectionDefinition"`
	LineStyleCollectionDefinition                       EntityListContainer `xml:"lineStyleCollectionDefinition"`
}

type InstrumentNameList struct {
	Entities EntityList `xml:"entities"`
	Language string     `xml:"language"`
}

type EntityListContainer struct {
	Entities EntityList `xml:"entities"`
}

type EntityList struct {
	IsArray  string          `xml:"array,attr"`
	Contents []ExpressionMap `xml:"ExpressionMapDefinition"`
}

type InitSwitchData struct {
	Enabled     bool       `xml:"enabled"`
	InitActions EntityList `xml:"initActions"`
}

type ExpressionMap struct {
	Name                          string                          `xml:"name"`
	EntityId                      string                          `xml:"entityID"`
	ParentEntityId                string                          `xml:"parentEntityID"`
	InheritanceMask               string                          `xml:"inheritanceMask"`
	Creator                       string                          `xml:"creator"`
	Description                   string                          `xml:"description"`
	Version                       string                          `xml:"version"`
	PluginNames                   string                          `xml:"pluginNames"`
	AutoMutualExclusion           bool                            `xml:"autoMutualExclusion"`
	AllowMultipleNotesAtSamePitch bool                            `xml:"allowMultipleNotesAtSamePitch"`
	InitSwitchData                InitSwitchData                  `xml:"initSwitchData"`
	Combinations                  PlayingTechniqueCombinationList `xml:"playingTechniqueCombinations"`
	TechniqueAddOns               TechniqueAddOnList              `xml:"techniqueAddOns"`
	MutualExclusionGroups         MutexGroupList                  `xml:"mutualExclusionGroups"`
}

type TechniqueAddOnList struct {
	IsArray         string           `xml:"array,attr"`
	TechniqueAddOns []TechniqueAddOn `xml:"techniqueAddOn"`
}

type TechniqueAddOn struct {
	SwitchID         int                 `xml:"techAddOnSwitchID"`
	TechniqueIDs     string              `xml:"techniqueIDs"`
	Enabled          bool                `xml:"enabled"`
	SwitchOnActions  SwitchOnActionList  `xml:"switchOnActions"`
	SwitchOffActions SwitchOffActionList `xml:"switchOffActions"`
}

type MutexGroupList struct {
	IsArray               string                 `xml:"array,attr"`
	MutualExclusionGroups []*MutualExclusionGroup `xml:"mutualExclusionGroup"`
}
type MutualExclusionGroup struct {
	GroupId      string `xml:"groupID"`
	Name         string `xml:"name"`
	TechniqueIds string `xml:"techniqueIDs"`
}

type PlayingTechniqueCombinationList struct {
	IsArray string                        `xml:"array,attr"`
	Combos  []*PlayingTechniqueCombination `xml:"playingTechniqueCombination"`
}

type PlayingTechniqueCombination struct {
	TechniqueIDs     string              `xml:"techniqueIDs"`
	BaseSwitchID     int                 `xml:"baseSwitchID"`
	Enabled          bool                `xml:"enabled"`
	Flags            int                 `xml:"flags"`
	ConditionString  string              `xml:"conditionString"`
	VelocityRange    string              `xml:"velocityRange"`
	PitchRange       string              `xml:"pitchRange"`
	Transpose        int                 `xml:"transpose"`
	TicksBefore      int                 `xml:"ticksBefore"`
	VelocityFactor   string              `xml:"velocityFactor"`
	LengthFactor     string              `xml:"lengthFactor"`
	VolumeType       VolumeType          `xml:"volumeType"`
	AttackType       AttackType          `xml:"attackType"`
	SwitchOnActions  SwitchOnActionList  `xml:"switchOnActions"`
	SwitchOffActions SwitchOffActionList `xml:"switchOffActions"`
}

type VolumeType struct {
	Type   string `xml:"type"`
	Param1 string `xml:"param1"`
}

type VolumeType2 struct {
	Type   string `xml:"type"`
	Param1 string `xml:"param1"`
}

type SwitchOnActionList struct {
	IsArray         string         `xml:"array,attr"`
	SwitchOnActions []SwitchAction `xml:"switchOnAction"`
}

type SwitchAction struct {
	//	XMLName string `xml:"switchOnAction"`
	Type   string `xml:"type"`
	Param1 string `xml:"param1"`
	Param2 string `xml:"param2"`
}

type SwitchOffActionList struct {
	IsArray          string         `xml:"array,attr"`
	SwitchOffActions []SwitchAction `xml:"switchOffAction"`
}

type SwitchOffAction struct {
	Type   string `xml:"type"`
	Param1 string `xml:"param1"`
	Param2 string `xml:"param2"`
}

type AttackType struct {
	Type   string `xml:"type"`
	Param1 string `xml:"param1"`
}

// Convert data to XML and write it to 'out'
func WriteXml(x interface{}, out io.Writer) error {
	_, err := io.WriteString(out, xml.Header)
	if err != nil {
		return fmt.Errorf("failed to write XML header: %w", err)
	}
	encoder := xml.NewEncoder(out)
	encoder.Indent("", "\t")
	xmlErr := encoder.Encode(x)
	if xmlErr != nil {
		return fmt.Errorf("failed to write XML: %w", xmlErr)
	}
	return nil
}

func ReadXml(bytes []byte) (*ScoreLib, error) {
	result := &ScoreLib{}
	err := xml.Unmarshal(bytes, result)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshall xml: %w", err)
	}
	return result, nil
}
