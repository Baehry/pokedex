package main

type NamedAPIResourceList struct {
	Count int
	Next string
	Previous string
	Results []NamedAPIResource
}

type APIResourceList struct {
	Count int
	Next string
	Previous string
	Results []APIResource
}

type Berry struct {
	Id int
	Name string
	GrowthTime int `json:"growth_time"`
	MaxHarvest int `json:"max_harvest"`
	NaturalGiftPower int `json:"natural_gift_power"`
	Size int
	Smoothness int
	SoilDryness int `json:"soil_dryness"`
	Firmness NamedAPIResource
	Flavors []BerryFlavorMap
	item NamedAPIResource
	NaturalGiftType NamedAPIResource `json:"natural_gift_type"`
}

type BerryFlavorMap struct {
	Potency int
	Flavor NamedAPIResource
}

type BerryFirmness struct {
	Id int
	Name string
	Berries []NamedAPIResource
	Names []Name
}

type BerryFlavors struct {
	Id int
	Name string
	Berries []FlavorBerryMap
	ContestType NamedAPIResource `json:"contest_type"`
	names []Name
}

type FlavorBerryMap struct {
	Potency int
	Berry NamedAPIResource
}

type ContestType struct {
	Id int
	Name string
	BerryFlavor NamedAPIResource `json:"berry_flavor"`
	Names []ContestName
}

type ContestName struct {
	Name string
	Color string
	Language NamedAPIResource
}

type ContestEffect struct {
	Id int
	Appeal int
	Jam int
	EffectEntries []Effect `json:"effect_entries"`
	FlavorTextEntries []FlavorText `json:"FlavorText"`
}

type SuperContestEffect struct {
	Id int
	Appeal int
	FlavorTextEntries []FlavorText `json:"FlavorText"`
	Moves []NamedAPIResource
}

type EncounterMethod struct {
	Id int
	Name string
	Order int
	Names []Name
}

type EncounterCondition struct {
	Id int
	Name string
	Names []Name
	Values []NamedAPIResource
}

type EncounterConditionValue struct {
	Id int
	Name string
	Condition NamedAPIResource
	Names []Name
}

type EvolutionChain struct {
	Id int
	BabyTriggerItem NamedAPIResource `json:"baby_trigger_item"`
	Chain ChainLink
}

type ChainLink struct {
	IsBaby bool `json:"is_baby"`
	Species NamedAPIResource
	EvolutionDetails []EvolutionDetail `json:"evolution_details"`
	EvolvesTo []ChainLink `json:"evolves_to"`
}

type EvolutionDetail struct {
	Item NamedAPIResource
	Trigger NamedAPIResource
	Gender int
	HeldItem NamedAPIResource `json:"held_item"`
	KnownMove NamedAPIResource `json:"known_move"`
	KnownMoveType Type `json:"known_move_type"`
	Location NamedAPIResource
	MinLevel int `json:"min_level"`
	MinHappiness int `json:"min_happiness"`
	MinBeauty int `json:"min_beauty"`
	MinAffection int `json:"min_affection"`
	NeedsOverworldRain bool `json:"needs_overworld_rain"`
	PartySpecies NamedAPIResource `json:"party_species"`
	PartyType NamedAPIResource `json:"party_type"`
	RelativePhysicalStats int `json:"relative_physical_stats"`
	TimeOfDay string `json:"time_of_day"`
	TradeSpecies NamedAPIResource `json:"trade_species"`
	TurnUpsideDown bool `json:"turn_upside_down"`
}

type EvolutionTrigger struct {
	Id int
	Name string
	Names []Name
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type Generation struct {
	Id int
	Name string
	Abilities []NamedAPIResource
	Names []Name
	MainRegion NamedAPIResource `json:"main_region"`
	Moves []NamedAPIResource
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
	Types []NamedAPIResource
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

type Pokedex struct {
	Id int
	Name string
	IsMainSeries bool `json:"is_main_series"`
	Descriptions []Description
	names []Name
	PokemonEntries []PokemonEntry `json:"pokemon_entries"`
	Region NamedAPIResource
	VersionGroups []NamedAPIResource `json:"version_group"`
}

type PokemonEntry struct {
	EntryNumber int `json:"entry_number"`
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type Version struct {
	Id int
	Name string
	Names []Name
	VersionGroup NamedAPIResource `json:"version_group"`
}

type VersionGroup struct {
	Id int
	Name string
	Order int
	Generation NamedAPIResource
	MovesLearnMethods NamedAPIResource `json:"move_learn_methods"`
	Pokedexes NamedAPIResource
	Regions NamedAPIResource
	Versions NamedAPIResource
}

type Item struct {
	Id int
	Name string
	Cost int
	FlingPower int `json:"fling_power"`
	FlingEffect NamedAPIResource `json:"fling_effect"`
	Attributes []NamedAPIResource
	Category NamedAPIResource
	EffectEntries []VerboseEffect `json:"effect_entries"`
	FlavorTextEntries []VersionGroupFlavorText `json:"flavor_text_entries"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Names []Name
	Sprites ItemSprites
	HeldByPokemon []ItemHolderPokemon `json:"held_by_pokemon"`
	BabyTriggerFor APIResource `json:"baby_trigger_for"`
	Machines []MachineVersionDetail
}

type ItemSprites struct {
	Default string
}

type ItemHolderPokemon struct {
	Pokemon NamedAPIResource
	VersionDetails []ItemHolderPokemonVersionDetail `json:"version_details"`
}

type ItemHolderPokemonVersionDetail struct {
	Rarity int
	Version NamedAPIResource
}

type ItemAttribute struct {
	Id int
	Name string
	Items []NamedAPIResource
	Names []Name
	Descriptions []Description
}

type ItemCategory struct {
	Id int
	Name string
	Items []NamedAPIResource
	Names []Name
	Pocket NamedAPIResource
}

type ItemFlingEffect struct {
	Id int
	Name string
	EffectEntries []Effect `json:"effect_entries"`
	Items []NamedAPIResource
}

type ItemPocket struct {
	Id int
	Name string
	Categories []NamedAPIResource
	Names []Name
}

type Location struct {
	Id int
	Name string
	Region NamedAPIResource
	Names []Name
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Areas []NamedAPIResource
}

type LocationArea struct {
	Id int
	Name string
	GameIndex int `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location NamedAPIResource
	Names []Name
	PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource `json:"encounter_method"`
	VersionDetails []EncounterVersionDetails `json:"version_details"`
}

type EncounterVersionDetails struct {
	Rate int
	Version NamedAPIResource
}

type PokemonEncounter struct {
	Pokemon NamedAPIResource
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PalParkArea struct {
	Id int
	Name string
	Names []Name
	PokemonEncounters []PalParkEncounterSpecies `json:"pokemon_encounters"`
}

type PalParkEncounterSpecies struct {
	BaseScore int `json:"base_score"`
	Rate int
	PokemonSpecies NamedAPIResource `json:"pokemon_species`
} 

type Region struct {
	Id int
	Locations []NamedAPIResource
	Name string
	Names []Name
	MainGeneration NamedAPIResource `json:"main_generation"`
	Pokedexes []NamedAPIResource
	VersionGroups []NamedAPIResource `json:"version_groups"`
}

type Machine struct {
	Id int
	Item NamedAPIResource
	Move NamedAPIResource
	VersionGroup NamedAPIResource `json:"version_group"`
}
type Move struct {
	Id int
	Name string
	Accuracy int
	EffectChance int `json:"effect_chance"`
	Pp int
	Priority int
	Power int
	ContestCombos ContestComboSets `json:"contest_combos"`
	ContestType NamedAPIResource `json:"contest_type"`
	ContestEffect APIResource `json:"contest_effect"`
	DamageClass NamedAPIResource `json:"damage_class"`
	EffectEntries []VerboseEffect `json:"effect_entries"`
	EffectChanges []AbilityEffectChange `json:"effect_changes"`
	LearnedByPokemon []NamedAPIResource `json:"learned_by_pokemon"`
	FlavorTextEntries []MoveFlavorText `json:"flavor_text_entries"`
	Generation NamedAPIResource
	Machines []MachineVersionDetail
	Meta MoveMetaData
	Names []Name
	PastValues []PastMoveStatValues `json:"past_values"`
	StatChanges []MoveStatChange `json:"stat_changes"`
	SuperContestEffect APIResource `json:"super_contest_effect"`
	Target NamedAPIResource
	Type NamedAPIResource
}

type ContestComboSets struct {
	Normal ContestComboDetail
	Super ContestComboDetail
}

type ContestComboDetail struct {
	UseBefore []NamedAPIResource `json:"use_before"`
	UseAfter []NamedAPIResource `json:"use_after"`
}

type MoveFlavorText struct {
	FlavorText string `json:"flavor_text"`
	Language NamedAPIResource
	VersionGroup NamedAPIResource
}

type MoveMetaData struct {
	Ailment NamedAPIResource
	Category NamedAPIResource
	MinHits int `json:"min_hits"`
	MaxHits int `json:"max_hits"`
	MinTurns int `json:"min_turns"`
	MaxTurns int `json:"max_turns"`
	Drain int
	Healing int
	CritRate int `json:"crit_rate"`
	AilmentChance int `json:"ailment_chance"`
	FlinchChance int `json:"flinch_chance"`
	StatChance int `json:"stat_chance"`
}

type MoveStatChange struct {
	Change int
	Stat NamedAPIResource
}

type PastMoveStatValues struct {
	Accuracy int
	EffectChance int `json:"effect_chance"`
	Power int
	Pp int
	EffectEntries []VerboseEffect `json:"effect_entries"`
	Type NamedAPIResource
	VersionGroup NamedAPIResource `json:"version_group"`
}

type MoveAilment struct {
	Id int
	Name string
	Moves []NamedAPIResource
	Names []Name
}

type MoveBattleStyle struct {
	Id int
	Name string
	Names []Name
}

type MoveCategory struct {
	Id int
	Name string
	Moves []NamedAPIResource
	Descriptions []Description
}

type MoveDamageClass struct {
	Id int
	Name string
	Descriptions []Description
	Moves []NamedAPIResource
	Names []Name
}

type MoveLearnMethod struct {
	Id int
	Name string
	Descriptions []Description
	Names []Name
	VersionGroups []NamedAPIResource `json:"version_group"`
}

type MoveTarget struct {
	Id int
	Name string
	Descriptions []Description
	Moves []NamedAPIResource
	Names []Name
}

type Ability struct {
	Id int
	Name string
	IsMainSeries bool `json:"is_main_series"`
	Generation NamedAPIResource
	Names []Name
	EffectEntries []VerboseEffect `json:"effect_entries"`
	EffectChanges []AbilityEffectChange `json:"effect_changes"`
	FlavorTextEntries []AbilityFlavorText `json:"flavor_text_entries"`
	Pokemon []AbilityPokemon
}

type AbilityEffectChange struct {
	EffectEntries []Effect `json:"effect_entries"`
	VersionGroup NamedAPIResource `json:"version_group"`
}

type AbilityFlavorText struct {
	FlavorText string `json:"flavor_text"`
	Language NamedAPIResource
	VersionGroup NamedAPIResource `json:"version_group"`
}

type AbilityPokemon struct {
	IsHidden bool `json:"is_hidden"`
	Slot int
	Pokemon NamedAPIResource
}

type Characteristic struct {
	Id int
	GeneModulo int `json:"gene_modulo"`
	PossibleValues []int `json:"possible_values"`
	HighestStat NamedAPIResource `json:"highest_stat"`
	Descriptions []Description
}

type EggGroup struct {
	Id int
	Name string
	Names []Name
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type Gender struct {
	Id int
	Name string
	PokemonSpeciesDetails []PokemonSpeciesGender `json:"pokemon_species_details"`
	RequiredForEvolution []NamedAPIResource `json:"required_for_evolution"`
}

type PokemonSpeciesGender struct {
	Rate int
	PokemonSpecies NamedAPIResource `json:"pokemon_species"`
}

type GrowthRate struct {
	Id int
	Name string
	Formula string
	Descriptions []Description
	Levels []GrowthRateExperienceLevel
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type GrowthRateExperienceLevel struct {
	Level int
	Experience int
}

type Nature struct {
	Id int
	Name string
	DecreasedStat NamedAPIResource `json:"decreased_stat"`
	IncreasedStat NamedAPIResource `json:"increased_stat"`
	HatesFlavor NamedAPIResource `json:"hates_flavor"`
	LikesFlavor NamedAPIResource `json:"likes_flavor"`
	PokeathlonStatChanges []NatureStatChange `json:"pokeathlon_stat_changes"`
	MoveBattleStylePreferences []MoveBattleStylePreference `json:"move_battle_style_preferences"`
	Names []Name
}

type NatureStatChange struct {
	MaxChange int `json:"max_change"`
	PokeathlonStat NamedAPIResource `json:"pokeathlon_stat"`
}

type MoveBattleStylePreference struct {
	LowHpPreference int `json:"low_hp_preference"`
	HighHpPreference int `json:"high_hp_preference"`
	MoveBattleStyle NamedAPIResource `json:"move_battle_style"`
}

type PokeathlonStat struct {
	Id int
	Name string
	Names []Name
	AffectingNatures NaturePokeathlonStatAffectSets `json:"affecting_natures"`
}

type NaturePokeathlonStatAffectSets struct {
	Increase []NaturePokeathlonStatAffect
	Decrease []NaturePokeathlonStatAffect
}

type NaturePokeathlonStatAffect struct {
	MaxChange int `json:"max_change"`
	nature NamedAPIResource
}

type Pokemon struct {
	Id int
	Name string
	BaseExperience int `json:"base_experience"`
	Height int
	IsDefault bool `json:"is_default"`
	Order int
	Weight int
	Abilities []PokemonAbility
	Forms []NamedAPIResource
	GameIndices []VersionGameIndex `json:"game_indices"`
	HeldItems []PokemonHeldItem `json:"held_items"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves []PokemonMove
	PastTypes []PokemonTypePast `json:"past_types"`
	PastAbilities []PokemonAbilityPast `json:"past_abilities"`
	Sprites PokemonSprites
	Cries PokemonCries
	Species NamedAPIResource
	Stats []PokemonStat
	Types []PokemonType
}

type PokemonAbility struct {
	IsHidden bool `json:"is_hidden"`
	Slot int
	Ability NamedAPIResource
}

type PokemonType struct {
	Slot int
	Type NamedAPIResource
}

type PokemonFormType struct {
	Slot int
	Type NamedAPIResource
}

type PokemonTypePast struct {
	Generation NamedAPIResource
	Types []PokemonType
}

type PokemonAbilityPast struct {
	Generation NamedAPIResource
	Abilities []PokemonAbility
}

type PokemonHeldItem struct {
	Item NamedAPIResource
	VersionDetails []PokemonHeldItemVersion `json:"version_details"`
}

type PokemonHeldItemVersion struct {
	Version NamedAPIResource
	Rarity int
}

type PokemonMove struct {
	Move NamedAPIResource
	VersionGroupDetails []PokemonMoveVersion `json:"version_group_details"`
}

type PokemonMoveVersion struct {
	MoveLearnMethod NamedAPIResource `json:"move_learn_method"`
	VersionGroup NamedAPIResource `json:"version_group"`
	LevelLearnedAt int `json:"level_learned_at"`
	order int
}

type PokemonStat struct {
	Stat NamedAPIResource
	Effort int
	BaseStat int `json:"base_stat"`
}

type PokemonSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
	FrontFemale string `json:"front_female"`
	FrontShinyFemale string `json:"front_shiny_female"`
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
	BackFemale string `json:"back_female"`
	BackShinyFemale string `json:"back_shiny_female"`
}

type PokemonCries struct {
	Latest string
	Legacy string
}

type LocationAreaEncounter struct {
	LocationArea NamedAPIResource `json:"location_area"`
	VersionDetails []VersionEncounterDetail `json:"version_details"`
}

type PokemonColor struct {
	Id int
	Name string
	Names []Name
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type PokemonForm struct {
	Id int
	Name string
	Order int
	FormOrder int `json:"form_order"`
	IsDefault bool `json:"is_default"`
	IsBattleOnly bool `json:"is_battle_only"`
	IsMega bool `json:"is_mega"`
	FormName string `json:"form_name"`
	Pokemon NamedAPIResource
	Types []PokemonFormType
	Sprites PokemonFormSprites
	VersionGroup NamedAPIResource `json:"version_group"`
	Names []Name
	FormNames []Name `json:"form_name"`
}

type PokemonFormSprites struct {
	FrontDefault string `json:"front_default"`
	FrontShiny string `json:"front_shiny"`
	BackDefault string `json:"back_default"`
	BackShiny string `json:"back_shiny"`
}

type PokemonHabitat struct {
	Id int
	Name string
	Names []Name
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type PokemonShape struct {
	Id int
	Name string
	AwesomeNames []AwesomeName `json:"awesome_names"`
	Names []Name
	PokemonSpecies []NamedAPIResource `json:"pokemon_species"`
}

type AwesomeName struct {
	AwesomeName string `json:"awesome_name"`
	Language NamedAPIResource
}

type PokemonSpecies struct {
	Id int
	Name string
	Order int
	GenderRate int `json:"gender_rate"`
	CaptureRate int `json:"capture_rate"`
	BaseHappiness int `json:"base_happiness"`
	IsBaby bool `json:"is_baby"`
	IsLegendary bool `json:"is_legendary"`
	IsMythical bool `json:"is_mythical"`
	HatchCounter int `json:"hatch_counter"`
	HasGenderDifferences bool `json:"has_gender_differences"`
	FormsSwitchable bool `json:"forms_switchable"`
	GrowthRate NamedAPIResource `json:"growth_rate"`
	PokedexNumbers []PokemonSpeciesDexEntry `json:"pokedex_numbers"`
	EggGroups []NamedAPIResource `json:"egg_groups"`
	Color NamedAPIResource
	Shape NamedAPIResource
	EvolvesFromSpecies NamedAPIResource `json:"evolves_from_species"`
	EvolutionChain APIResource `json:"evolution_chain"`
	Habitat NamedAPIResource
	Generation NamedAPIResource
	Names []Name
	PalParkEncounters []PalPArkEncounterArea `json:"pa_park_encounters"`
	FlavorTextEntries []FlavorText `json:"flavor_text_entries"`
	FormDesriptions []Description `json:"form_descriptions"`
	Genera []Genus
	Varieties []PokemonSpeciesVariety
}

type Genus struct {
	Genus string
	Language NamedAPIResource
}

type PokemonSpeciesDexEntry struct {
	EntryNumber int `json:"entry_number"`
	Pokedex NamedAPIResource
}

type PalPArkEncounterArea struct {
	BaseScore int `json:"base_score"`
	Rate int
	Area NamedAPIResource
}

type PokemonSpeciesVariety struct {
	IsDefault bool `json:"is_default"`
	Pokemon NamedAPIResource
}

type Stat struct {
	Id int
	Name string
	GameIndex int `json:"game_index"`
	IsBattleOnly bool `json:"is_battle_only"`
	AffectingMoves MoveStatAffectSets `json:"affecting_moves"`
	AffectingNatures NatureStatAffectSets `json:"affecting_natures"`
	Characteristics []APIResource
	MoveDamageClass NamedAPIResource `json:"move_damage_class"`
	Names []Name
}

type MoveStatAffectSets struct {
	Increase []MoveStatAffect
	Decrease []MoveStatAffect
}

type MoveStatAffect struct {
	Change int
	Move NamedAPIResource
}

type NatureStatAffectSets struct {
	Increase []NamedAPIResource
	Decrease []NamedAPIResource
}

type Type struct {
	Id int
	Name string
	DamageRelations TypeRelations `json:"damage_relations"`
	PastDamageRelations []TypeRelationsPast `json:"past_damage_relations"`
	GameIndices []GenerationGameIndex `json:"game_indices"`
	Generation NamedAPIResource
	MoveDamageClass NamedAPIResource `json:"move_damage_class"`
	Names []Name
	Pokemon []TypePokemon
	Moves []NamedAPIResource
}

type TypePokemon struct {
	Slot int
	Pokemon NamedAPIResource
}

type TypeRelations struct {
	NoDamageTo []NamedAPIResource `json:"no_damage_to"`
	HalfDamageTo []NamedAPIResource `json:"half_damage_to"`
	DoubleDamageTo []NamedAPIResource `json:"double_damage_to"`
	NoDamageFrom []NamedAPIResource `json:"no_damage_from"`
	HalfDamageFrom []NamedAPIResource `json:"half_damage_from"`
	DoubleDamageFrom []NamedAPIResource `json:"double_damage_from"`
}

type TypeRelationsPast struct {
	Generation NamedAPIResource
	DamageRelations TypeRelations `json:"damage_relations"`
}

type Language struct {
	Id int
	Name string
	Official bool
	Iso639 string
	Iso3166 string
	Names []Name
}

type APIResource struct {
	Url string
}

type Description struct {
	Description string
	Language NamedAPIResource
}

type Effect struct {
	Effect string
	Language NamedAPIResource
}

type Encounter struct {
	MinLevel int `json:"min_level"`
	MaxLevel int `json:"max_level"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	Chance int
	Method NamedAPIResource
}

type FlavorText struct {
	FlavorText string `json:"flavor_text"`
	Language NamedAPIResource
	Version NamedAPIResource
}

type GenerationGameIndex struct {
	GameIndex int `json:"game_index"`
	Generation NamedAPIResource
}

type MachineVersionDetail struct {
	Machine APIResource
	VersionGroup NamedAPIResource `json:"version_group"`
}

type Name struct {
	Name string
	Language NamedAPIResource
}

type NamedAPIResource struct {
	Name string
	Url string
}

type VerboseEffect struct {
	Effect string
	ShortEffect string `json:"short_effect"`
	Language NamedAPIResource
}

type VersionEncounterDetail struct {
	Version NamedAPIResource
	MaxChance int `json:"max_chance"`
	EncounterDetails []Encounter `json:"encounter_details"`
}

type VersionGameIndex struct {
	GameIndex int `json:"game_index"`
	Version NamedAPIResource
}

type VersionGroupFlavorText struct {
	Text string
	Language NamedAPIResource
	VersionGroup NamedAPIResource `json:"version_group"`
}