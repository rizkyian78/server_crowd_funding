package campaign

type Service interface {
	FindCampaigns(ID string) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(ID string) ([]Campaign, error) {
	if ID != "" {
		campaigns, err := s.repository.FindByUserID(ID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}
	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
