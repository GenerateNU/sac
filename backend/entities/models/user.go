package models

import (
	"gorm.io/gorm"
)

type UserRole string

var (
	Super   UserRole = "super"
	Student UserRole = "student"
)

type College string

const (
	CAMD College = "CAMD" // College of Arts, Media and Design
	DMSB College = "DMSB" // D'Amore-McKim School of Business
	KCCS College = "KCCS" // Khoury College of Computer Sciences
	CE   College = "CE"   // College of Engineering
	BCHS College = "BCHS" // Bouvé College of Health Sciences
	SL   College = "SL"   // School of Law
	CPS  College = "CPS"  // College of Professional Studies
	CS   College = "CS"   // College of Science
	CSSH College = "CSSH" // College of Social Sciences and Humanities
)

type Major string

// see https://admissions.northeastern.edu/academics/areas-of-study/
const (
	AfricanaStudies                                           Major = "africanaStudies"
	AmericanSignLanguage                                      Major = "americanSignLanguage"
	AmericanSignLanguageEnglishInterpreting                   Major = "americanSignLanguage-EnglishInterpreting"
	AppliedPhysics                                            Major = "appliedPhysics"
	ArchitecturalStudies                                      Major = "architecturalStudies"
	Architecture                                              Major = "architecture"
	ArtArtVisualStudies                                       Major = "art:ArtVisualStudies"
	BehavioralNeuroscience                                    Major = "behavioralNeuroscience"
	Biochemistry                                              Major = "biochemistry"
	Bioengineering                                            Major = "bioengineering"
	Biology                                                   Major = "biology"
	BiomedicalPhysics                                         Major = "biomedicalPhysics"
	BusinessAdministration                                    Major = "businessAdministration"
	BusinessAdministrationAccounting                          Major = "businessAdministration:Accounting"
	BusinessAdministrationAccountingAndAdvisoryServices       Major = "businessAdministration:AccountingAndAdvisoryServices"
	BusinessAdministrationBrandManagement                     Major = "businessAdministration:BrandManagement"
	BusinessAdministrationBusinessAnalytics                   Major = "businessAdministration:BusinessAnalytics"
	BusinessAdministrationCorporateInnovation                 Major = "businessAdministration:CorporateInnovation"
	BusinessAdministrationEntrepreneurialStartups             Major = "businessAdministration:EntrepreneurialStartups"
	BusinessAdministrationFamilyBusiness                      Major = "businessAdministration:FamilyBusiness"
	BusinessAdministrationFinance                             Major = "businessAdministration:Finance"
	BusinessAdministrationFintech                             Major = "businessAdministration:Fintech"
	BusinessAdministrationHealthcareManagementAndConsulting   Major = "businessAdministration:HealthcareManagementAndConsulting"
	BusinessAdministrationManagement                          Major = "businessAdministration:Management"
	BusinessAdministrationManagementInformationSystems        Major = "businessAdministration:ManagementInformationSystems"
	BusinessAdministrationMarketing                           Major = "businessAdministration:Marketing"
	BusinessAdministrationMarketingAnalytics                  Major = "businessAdministration:MarketingAnalytics"
	BusinessAdministrationSocialInnovationAndEntrepreneurship Major = "businessAdministration:SocialInnovationAndEntrepreneurship"
	BusinessAdministrationSupplyChainManagement               Major = "businessAdministration:SupplyChainManagement"
	CellAndMolecularBiology                                   Major = "cellAndMolecularBiology"
	ChemicalEngineering                                       Major = "chemicalEngineering"
	Chemistry                                                 Major = "chemistry"
	CivilEngineering                                          Major = "civilEngineering"
	CommunicationStudies                                      Major = "communicationStudies"
	ComputerEngineering                                       Major = "computerEngineering"
	ComputerScience                                           Major = "computerScience"
	ComputingAndLaw                                           Major = "computingAndLaw"
	CriminologyAndCriminalJustice                             Major = "criminologyAndCriminalJustice"
	CulturalAnthropology                                      Major = "culturalAnthropology"
	Cybersecurity                                             Major = "cybersecurity"
	DataScience                                               Major = "dataScience"
	Design                                                    Major = "design"
	Economics                                                 Major = "economics"
	ElectricalEngineering                                     Major = "electricalEngineering"
	English                                                   Major = "english"
	EnvironmentalAndSustainabilityStudies                     Major = "environmentalAndSustainabilityStudies"
	EnvironmentalEngineering                                  Major = "environmentalEngineering"
	EnvironmentalScience                                      Major = "environmentalScience"
	EnvironmentalStudies                                      Major = "environmentalStudies"
	GameArtAndAnimation                                       Major = "gameArtAndAnimation"
	GameDesign                                                Major = "gameDesign"
	GlobalAsianStudies                                        Major = "globalAsianStudies"
	HealthScience                                             Major = "healthScience"
	History                                                   Major = "history"
	HistoryCultureAndLaw                                      Major = "historyCultureAndLaw"
	HumanServices                                             Major = "humanServices"
	IndustrialEngineering                                     Major = "industrialEngineering"
	InternationalAffairs                                      Major = "internationalAffairs"
	InternationalBusiness                                     Major = "internationalBusiness"
	InternationalBusinessAccounting                           Major = "internationalBusiness:Accounting"
	InternationalBusinessAccountingAndAdvisoryServices        Major = "internationalBusiness:AccountingAndAdvisoryServices"
	InternationalBusinessBrandManagement                      Major = "internationalBusiness:BrandManagement"
	InternationalBusinessBusinessAnalytics                    Major = "internationalBusiness:BusinessAnalytics"
	InternationalBusinessCorporateInnovation                  Major = "internationalBusiness:CorporateInnovation"
	InternationalBusinessEntrepreneurialStartups              Major = "internationalBusiness:EntrepreneurialStartups"
	InternationalBusinessFamilyBusiness                       Major = "internationalBusiness:FamilyBusiness"
	InternationalBusinessFinance                              Major = "internationalBusiness:Finance"
	InternationalBusinessFintech                              Major = "internationalBusiness:Fintech"
	InternationalBusinessHealthcareManagementAndConsulting    Major = "internationalBusiness:HealthcareManagementAndConsulting"
	InternationalBusinessManagement                           Major = "internationalBusiness:Management"
	InternationalBusinessManagementInformationSystems         Major = "internationalBusiness:ManagementInformationSystems"
	InternationalBusinessMarketing                            Major = "internationalBusiness:Marketing"
	InternationalBusinessMarketingAnalytics                   Major = "internationalBusiness:MarketingAnalytics"
	InternationalBusinessSocialInnovationAndEntrepreneurship  Major = "internationalBusiness:SocialInnovationAndEntrepreneurship"
	InternationalBusinessSupplyChainManagement                Major = "internationalBusiness:SupplyChainManagement"
	Journalism                                                Major = "journalism"
	LandscapeArchitecture                                     Major = "landscapeArchitecture"
	Linguistics                                               Major = "linguistics"
	MarineBiology                                             Major = "marineBiology"
	Mathematics                                               Major = "mathematics"
	MechanicalEngineering                                     Major = "mechanicalEngineering"
	MediaAndScreenStudies                                     Major = "mediaAndScreenStudies"
	MediaArts                                                 Major = "mediaArts"
	Music                                                     Major = "music"
	MusicTechnology                                           Major = "musicTechnology"
	Nursing                                                   Major = "nursing"
	PharmaceuticalSciences                                    Major = "pharmaceuticalSciences"
	PharmacyPharmD                                            Major = "pharmacy(PharmD)"
	Philosophy                                                Major = "philosophy"
	Physics                                                   Major = "physics"
	PoliticalScience                                          Major = "politicalScience"
	PoliticsPhilosophyEconomics                               Major = "politicsPhilosophyEconomics"
	Psychology                                                Major = "psychology"
	PublicHealth                                              Major = "publicHealth"
	PublicRelations                                           Major = "publicRelations"
	ReligiousStudies                                          Major = "religiousStudies"
	Sociology                                                 Major = "sociology"
	Spanish                                                   Major = "spanish"
	SpeechLanguagePathologyAndAudiology                       Major = "speechLanguagePathologyAndAudiology"
	Theatre                                                   Major = "theatre"
)

type GraduationCycle string

const (
	December GraduationCycle = "december"
	May      GraduationCycle = "may"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type User struct {
	Model

	Role            UserRole        `json:"role"`
	FirstName       string          `json:"first_name"`
	LastName        string          `json:"last_name"`
	Email           string          `json:"email"`
	PasswordHash    string          `json:"-"`
	Major0          Major           `json:"major0"`
	Major1          Major           `json:"major1"`
	Major2          Major           `json:"major2"`
	College         College         `json:"college"`
	GraduationCycle GraduationCycle `json:"graduation_cycle"`
	GraduationYear  int16           `json:"graduation_year"`
	IsVerified      bool            `json:"is_verified"`

	Tag               []Tag     `gorm:"many2many:user_tags;" json:"-"`
	Admin             []Club    `gorm:"many2many:user_club_admins;" json:"-"`
	Member            []Club    `gorm:"many2many:user_club_members;" json:"-"`
	Follower          []Club    `gorm:"many2many:user_club_followers;" json:"-"`
	IntendedApplicant []Club    `gorm:"many2many:user_club_intended_applicants;" json:"-"`
	Asked             []Comment `gorm:"foreignKey:AskedByID;" json:"-"`
	Answered          []Comment `gorm:"foreignKey:AnsweredByID;" json:"-"`
	RSVP              []Event   `gorm:"many2many:user_event_rsvps;" json:"-"`
	Waitlist          []Event   `gorm:"many2many:user_event_waitlists;" json:"-"`
}

type CreateUserRequestBody struct {
	FirstName string `json:"first_name" validate:"required,max=255"`
	LastName  string `json:"last_name" validate:"required,max=255"`
	Email     string `json:"email" validate:"required,email,neu_email,max=255"`
	Password  string `json:"password" validate:"required,max=255"` // MARK: must be validated manually
	// Optional fields
	Major0          Major           `json:"major0" validate:"not_equal_if_not_empty=Major1,not_equal_if_not_empty=Major2,omitempty,max=255,oneof=africanaStudies americanSignLanguage americanSignLanguage-EnglishInterpreting appliedPhysics architecturalStudies architecture art:ArtVisualStudies behavioralNeuroscience biochemistry bioengineering biology biomedicalPhysics businessAdministration businessAdministration:Accounting businessAdministration:AccountingAndAdvisoryServices businessAdministration:BrandManagement businessAdministration:BusinessAnalytics businessAdministration:CorporateInnovation businessAdministration:EntrepreneurialStartups businessAdministration:FamilyBusiness businessAdministration:Finance businessAdministration:Fintech businessAdministration:HealthcareManagementAndConsulting businessAdministration:Management businessAdministration:ManagementInformationSystems businessAdministration:Marketing businessAdministration:MarketingAnalytics businessAdministration:SocialInnovationAndEntrepreneurship businessAdministration:SupplyChainManagement cellAndMolecularBiology chemicalEngineering chemistry civilEngineering communicationStudies computerEngineering computerScience computingAndLaw criminologyAndCriminalJustice culturalAnthropology cybersecurity dataScience design economics electricalEngineering english environmentalAndSustainabilityStudies environmentalEngineering environmentalScience environmentalStudies gameArtAndAnimation gameDesign globalAsianStudies healthScience history historyCultureAndLaw humanServices industrialEngineering internationalAffairs internationalBusiness internationalBusiness:Accounting internationalBusiness:AccountingAndAdvisoryServices internationalBusiness:BrandManagement internationalBusiness:BusinessAnalytics internationalBusiness:CorporateInnovation internationalBusiness:EntrepreneurialStartups internationalBusiness:FamilyBusiness internationalBusiness:Finance internationalBusiness:Fintech internationalBusiness:HealthcareManagementAndConsulting internationalBusiness:Management internationalBusiness:ManagementInformationSystems internationalBusiness:Marketing internationalBusiness:MarketingAnalytics internationalBusiness:SocialInnovationAndEntrepreneurship internationalBusiness:SupplyChainManagement journalism landscapeArchitecture linguistics marineBiology mathematics mechanicalEngineering mediaAndScreenStudies mediaArts music musicTechnology nursing pharmaceuticalSciences pharmacy(PharmD) philosophy physics politicalScience politicsPhilosophyEconomics psychology publicHealth publicRelations religiousStudies sociology spanish speechLanguagePathologyAndAudiology theatre"`
	Major1          Major           `json:"major1" validate:"not_equal_if_not_empty=Major0,not_equal_if_not_empty=Major2,omitempty,max=255,oneof=africanaStudies americanSignLanguage americanSignLanguage-EnglishInterpreting appliedPhysics architecturalStudies architecture art:ArtVisualStudies behavioralNeuroscience biochemistry bioengineering biology biomedicalPhysics businessAdministration businessAdministration:Accounting businessAdministration:AccountingAndAdvisoryServices businessAdministration:BrandManagement businessAdministration:BusinessAnalytics businessAdministration:CorporateInnovation businessAdministration:EntrepreneurialStartups businessAdministration:FamilyBusiness businessAdministration:Finance businessAdministration:Fintech businessAdministration:HealthcareManagementAndConsulting businessAdministration:Management businessAdministration:ManagementInformationSystems businessAdministration:Marketing businessAdministration:MarketingAnalytics businessAdministration:SocialInnovationAndEntrepreneurship businessAdministration:SupplyChainManagement cellAndMolecularBiology chemicalEngineering chemistry civilEngineering communicationStudies computerEngineering computerScience computingAndLaw criminologyAndCriminalJustice culturalAnthropology cybersecurity dataScience design economics electricalEngineering english environmentalAndSustainabilityStudies environmentalEngineering environmentalScience environmentalStudies gameArtAndAnimation gameDesign globalAsianStudies healthScience history historyCultureAndLaw humanServices industrialEngineering internationalAffairs internationalBusiness internationalBusiness:Accounting internationalBusiness:AccountingAndAdvisoryServices internationalBusiness:BrandManagement internationalBusiness:BusinessAnalytics internationalBusiness:CorporateInnovation internationalBusiness:EntrepreneurialStartups internationalBusiness:FamilyBusiness internationalBusiness:Finance internationalBusiness:Fintech internationalBusiness:HealthcareManagementAndConsulting internationalBusiness:Management internationalBusiness:ManagementInformationSystems internationalBusiness:Marketing internationalBusiness:MarketingAnalytics internationalBusiness:SocialInnovationAndEntrepreneurship internationalBusiness:SupplyChainManagement journalism landscapeArchitecture linguistics marineBiology mathematics mechanicalEngineering mediaAndScreenStudies mediaArts music musicTechnology nursing pharmaceuticalSciences pharmacy(PharmD) philosophy physics politicalScience politicsPhilosophyEconomics psychology publicHealth publicRelations religiousStudies sociology spanish speechLanguagePathologyAndAudiology theatre"`
	Major2          Major           `json:"major2" validate:"not_equal_if_not_empty=Major0,not_equal_if_not_empty=Major1,omitempty,max=255,oneof=africanaStudies americanSignLanguage americanSignLanguage-EnglishInterpreting appliedPhysics architecturalStudies architecture art:ArtVisualStudies behavioralNeuroscience biochemistry bioengineering biology biomedicalPhysics businessAdministration businessAdministration:Accounting businessAdministration:AccountingAndAdvisoryServices businessAdministration:BrandManagement businessAdministration:BusinessAnalytics businessAdministration:CorporateInnovation businessAdministration:EntrepreneurialStartups businessAdministration:FamilyBusiness businessAdministration:Finance businessAdministration:Fintech businessAdministration:HealthcareManagementAndConsulting businessAdministration:Management businessAdministration:ManagementInformationSystems businessAdministration:Marketing businessAdministration:MarketingAnalytics businessAdministration:SocialInnovationAndEntrepreneurship businessAdministration:SupplyChainManagement cellAndMolecularBiology chemicalEngineering chemistry civilEngineering communicationStudies computerEngineering computerScience computingAndLaw criminologyAndCriminalJustice culturalAnthropology cybersecurity dataScience design economics electricalEngineering english environmentalAndSustainabilityStudies environmentalEngineering environmentalScience environmentalStudies gameArtAndAnimation gameDesign globalAsianStudies healthScience history historyCultureAndLaw humanServices industrialEngineering internationalAffairs internationalBusiness internationalBusiness:Accounting internationalBusiness:AccountingAndAdvisoryServices internationalBusiness:BrandManagement internationalBusiness:BusinessAnalytics internationalBusiness:CorporateInnovation internationalBusiness:EntrepreneurialStartups internationalBusiness:FamilyBusiness internationalBusiness:Finance internationalBusiness:Fintech internationalBusiness:HealthcareManagementAndConsulting internationalBusiness:Management internationalBusiness:ManagementInformationSystems internationalBusiness:Marketing internationalBusiness:MarketingAnalytics internationalBusiness:SocialInnovationAndEntrepreneurship internationalBusiness:SupplyChainManagement journalism landscapeArchitecture linguistics marineBiology mathematics mechanicalEngineering mediaAndScreenStudies mediaArts music musicTechnology nursing pharmaceuticalSciences pharmacy(PharmD) philosophy physics politicalScience politicsPhilosophyEconomics psychology publicHealth publicRelations religiousStudies sociology spanish speechLanguagePathologyAndAudiology theatre"`
	College         College         `json:"college" validate:"omitempty,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	GraduationCycle GraduationCycle `json:"graduation_cycle" validate:"omitempty,max=255,oneof=december may"`
	GraduationYear  int16           `json:"graduation_year" validate:"omitempty"`
}

type UpdateUserRequestBody struct {
	FirstName       string          `json:"first_name" validate:"omitempty,max=255"`
	LastName        string          `json:"last_name" validate:"omitempty,max=255"`
	Major0          Major           `json:"major0" validate:"not_equal_if_not_empty=Major1,not_equal_if_not_empty=Major2,omitempty,max=255,oneof=africanaStudies americanSignLanguage americanSignLanguage-EnglishInterpreting appliedPhysics architecturalStudies architecture art:ArtVisualStudies behavioralNeuroscience biochemistry bioengineering biology biomedicalPhysics businessAdministration businessAdministration:Accounting businessAdministration:AccountingAndAdvisoryServices businessAdministration:BrandManagement businessAdministration:BusinessAnalytics businessAdministration:CorporateInnovation businessAdministration:EntrepreneurialStartups businessAdministration:FamilyBusiness businessAdministration:Finance businessAdministration:Fintech businessAdministration:HealthcareManagementAndConsulting businessAdministration:Management businessAdministration:ManagementInformationSystems businessAdministration:Marketing businessAdministration:MarketingAnalytics businessAdministration:SocialInnovationAndEntrepreneurship businessAdministration:SupplyChainManagement cellAndMolecularBiology chemicalEngineering chemistry civilEngineering communicationStudies computerEngineering computerScience computingAndLaw criminologyAndCriminalJustice culturalAnthropology cybersecurity dataScience design economics electricalEngineering english environmentalAndSustainabilityStudies environmentalEngineering environmentalScience environmentalStudies gameArtAndAnimation gameDesign globalAsianStudies healthScience history historyCultureAndLaw humanServices industrialEngineering internationalAffairs internationalBusiness internationalBusiness:Accounting internationalBusiness:AccountingAndAdvisoryServices internationalBusiness:BrandManagement internationalBusiness:BusinessAnalytics internationalBusiness:CorporateInnovation internationalBusiness:EntrepreneurialStartups internationalBusiness:FamilyBusiness internationalBusiness:Finance internationalBusiness:Fintech internationalBusiness:HealthcareManagementAndConsulting internationalBusiness:Management internationalBusiness:ManagementInformationSystems internationalBusiness:Marketing internationalBusiness:MarketingAnalytics internationalBusiness:SocialInnovationAndEntrepreneurship internationalBusiness:SupplyChainManagement journalism landscapeArchitecture linguistics marineBiology mathematics mechanicalEngineering mediaAndScreenStudies mediaArts music musicTechnology nursing pharmaceuticalSciences pharmacy(PharmD) philosophy physics politicalScience politicsPhilosophyEconomics psychology publicHealth publicRelations religiousStudies sociology spanish speechLanguagePathologyAndAudiology theatre"`
	Major1          Major           `json:"major1" validate:"not_equal_if_not_empty=Major0,not_equal_if_not_empty=Major2,omitempty,max=255,oneof=africanaStudies americanSignLanguage americanSignLanguage-EnglishInterpreting appliedPhysics architecturalStudies architecture art:ArtVisualStudies behavioralNeuroscience biochemistry bioengineering biology biomedicalPhysics businessAdministration businessAdministration:Accounting businessAdministration:AccountingAndAdvisoryServices businessAdministration:BrandManagement businessAdministration:BusinessAnalytics businessAdministration:CorporateInnovation businessAdministration:EntrepreneurialStartups businessAdministration:FamilyBusiness businessAdministration:Finance businessAdministration:Fintech businessAdministration:HealthcareManagementAndConsulting businessAdministration:Management businessAdministration:ManagementInformationSystems businessAdministration:Marketing businessAdministration:MarketingAnalytics businessAdministration:SocialInnovationAndEntrepreneurship businessAdministration:SupplyChainManagement cellAndMolecularBiology chemicalEngineering chemistry civilEngineering communicationStudies computerEngineering computerScience computingAndLaw criminologyAndCriminalJustice culturalAnthropology cybersecurity dataScience design economics electricalEngineering english environmentalAndSustainabilityStudies environmentalEngineering environmentalScience environmentalStudies gameArtAndAnimation gameDesign globalAsianStudies healthScience history historyCultureAndLaw humanServices industrialEngineering internationalAffairs internationalBusiness internationalBusiness:Accounting internationalBusiness:AccountingAndAdvisoryServices internationalBusiness:BrandManagement internationalBusiness:BusinessAnalytics internationalBusiness:CorporateInnovation internationalBusiness:EntrepreneurialStartups internationalBusiness:FamilyBusiness internationalBusiness:Finance internationalBusiness:Fintech internationalBusiness:HealthcareManagementAndConsulting internationalBusiness:Management internationalBusiness:ManagementInformationSystems internationalBusiness:Marketing internationalBusiness:MarketingAnalytics internationalBusiness:SocialInnovationAndEntrepreneurship internationalBusiness:SupplyChainManagement journalism landscapeArchitecture linguistics marineBiology mathematics mechanicalEngineering mediaAndScreenStudies mediaArts music musicTechnology nursing pharmaceuticalSciences pharmacy(PharmD) philosophy physics politicalScience politicsPhilosophyEconomics psychology publicHealth publicRelations religiousStudies sociology spanish speechLanguagePathologyAndAudiology theatre"`
	Major2          Major           `json:"major2" validate:"not_equal_if_not_empty=Major0,not_equal_if_not_empty=Major1,omitempty,max=255,oneof=africanaStudies americanSignLanguage americanSignLanguage-EnglishInterpreting appliedPhysics architecturalStudies architecture art:ArtVisualStudies behavioralNeuroscience biochemistry bioengineering biology biomedicalPhysics businessAdministration businessAdministration:Accounting businessAdministration:AccountingAndAdvisoryServices businessAdministration:BrandManagement businessAdministration:BusinessAnalytics businessAdministration:CorporateInnovation businessAdministration:EntrepreneurialStartups businessAdministration:FamilyBusiness businessAdministration:Finance businessAdministration:Fintech businessAdministration:HealthcareManagementAndConsulting businessAdministration:Management businessAdministration:ManagementInformationSystems businessAdministration:Marketing businessAdministration:MarketingAnalytics businessAdministration:SocialInnovationAndEntrepreneurship businessAdministration:SupplyChainManagement cellAndMolecularBiology chemicalEngineering chemistry civilEngineering communicationStudies computerEngineering computerScience computingAndLaw criminologyAndCriminalJustice culturalAnthropology cybersecurity dataScience design economics electricalEngineering english environmentalAndSustainabilityStudies environmentalEngineering environmentalScience environmentalStudies gameArtAndAnimation gameDesign globalAsianStudies healthScience history historyCultureAndLaw humanServices industrialEngineering internationalAffairs internationalBusiness internationalBusiness:Accounting internationalBusiness:AccountingAndAdvisoryServices internationalBusiness:BrandManagement internationalBusiness:BusinessAnalytics internationalBusiness:CorporateInnovation internationalBusiness:EntrepreneurialStartups internationalBusiness:FamilyBusiness internationalBusiness:Finance internationalBusiness:Fintech internationalBusiness:HealthcareManagementAndConsulting internationalBusiness:Management internationalBusiness:ManagementInformationSystems internationalBusiness:Marketing internationalBusiness:MarketingAnalytics internationalBusiness:SocialInnovationAndEntrepreneurship internationalBusiness:SupplyChainManagement journalism landscapeArchitecture linguistics marineBiology mathematics mechanicalEngineering mediaAndScreenStudies mediaArts music musicTechnology nursing pharmaceuticalSciences pharmacy(PharmD) philosophy physics politicalScience politicsPhilosophyEconomics psychology publicHealth publicRelations religiousStudies sociology spanish speechLanguagePathologyAndAudiology theatre"`
	College         College         `json:"college" validate:"omitempty,oneof=CAMD DMSB KCCS CE BCHS SL CPS CS CSSH"`
	GraduationCycle GraduationCycle `json:"graduation_cycle" validate:"omitempty,max=255,oneof=december may"`
	GraduationYear  int16           `json:"graduation_year" validate:"omitempty"`
}

type LoginUserResponseBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,max=255"` // MARK: must be validated manually
}

type UpdatePasswordRequestBody struct {
	OldPassword string `json:"old_password" validate:"required,max=255"`                                    // MARK: must be validated manually
	NewPassword string `json:"new_password" validate:"required,not_equal_if_not_empty=OldPassword,max=255"` // MARK: must be validated manually
}

type RefreshTokenRequestBody struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	sac := &Club{}
	if err := tx.Where("name = ?", "SAC").First(sac).Error; err != nil {
		return err
	}

	if err := tx.Model(u).Association("Member").Append(sac); err != nil {
		return err
	}

	if err := tx.Model(u).Association("Follower").Append(sac); err != nil {
		return err
	}

	if err := tx.Model(&Club{}).Where("id = ?", sac.ID).Update("num_members", gorm.Expr("num_members + 1")).Error; err != nil {
		return err
	}

	return nil
}

func (u *User) AfterDelete(tx *gorm.DB) (err error) {
	sac := &Club{}
	if err := tx.Where("name = ?", "SAC").First(sac).Error; err != nil {
		return err
	}

	if err := tx.Model(&Club{}).Where("id = ?", sac.ID).Update("num_members", gorm.Expr("num_members - 1")).Error; err != nil {
		return err
	}

	return nil
}
