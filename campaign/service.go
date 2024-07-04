package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId int) ([]Campaign, error)
	GetCampaignById(input GetCampaignDetailInput) (Campaign, error)
	CreateCampign(input CreateCampignInput) (Campaign, error)
	UpdateCampaign(inputID GetCampaignDetailInput, inputData UpdateCampaignInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId int) ([]Campaign, error) {
	if userId != 0 {
		campaigns, err := s.repository.FindByUserID(userId)
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

func (s *service) GetCampaignById(input GetCampaignDetailInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(input.ID)

	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) CreateCampign(input CreateCampignInput) (Campaign, error) {

	campaign := Campaign{}

	campaign.Name = input.Name
	campaign.ShortDescription = input.ShortDescription
	campaign.Description = input.Description
	campaign.Perks = input.Perks
	campaign.GoalAmount = input.GoalAmount
	campaign.UserID = input.User.ID

	slugCandidate := fmt.Sprintf("%s %d", input.Name, input.User.ID)
	campaign.Slug = slug.Make(slugCandidate)

	// pembuatan slug
	newCampign, err := s.repository.Save(campaign)
	if err != nil {
		return newCampign, err
	}

	return newCampign, nil
}

func (s *service) UpdateCampaign(inputID GetCampaignDetailInput, inputData UpdateCampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindByID(inputID.ID)
	if err != nil {
		return campaign, err
	}

	if campaign.UserID != inputData.User.ID {
		return campaign, errors.New("not an owner of the campaign")
	}

	if inputData.Name != "" {
		campaign.Name = inputData.Name
	}

	if inputData.ShortDescription != "" {
		campaign.ShortDescription = inputData.ShortDescription
	}

	if inputData.Description != "" {
		campaign.Description = inputData.Description
	}

	if inputData.Perks != "" {
		campaign.Perks = inputData.Perks
	}

	if inputData.GoalAmount != 0 {
		campaign.GoalAmount = inputData.GoalAmount
	}

	if inputData.User.ID != 0 {
		campaign.UserID = inputData.User.ID
		slugCandidate := fmt.Sprintf("%s %d", inputData.Name, inputData.User.ID)
		campaign.Slug = slug.Make(slugCandidate)
	}

	updatedCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return updatedCampaign, err
	}

	return updatedCampaign, nil
}
