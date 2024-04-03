package repo

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/pjotrscholtze/dmailserver-rest-api/models"
)

type setupRepo struct {
	commandPrefix string
}
type SetupRepo interface {
	ListEmail() ([]*models.EmailAccountListItem, error)
	AddEmail(account models.EmailAccount) error
	UpdateEmail(account models.EmailAccount) error
	RemoveEmail(emailAddress string) error
	HasEmail(emailAddress string) (bool, error)

	GetEmail(addres string) (*models.EmailAccountListItem, error)
	HasAlias(aliasAddress, recepientAddres string) (bool, error)
	AddAlias(aliasAddress, recepientAddres string) error
	RemoveAlias(aliasAddress, recepientAddres string) error

	AddressExistsOnServer(addres string) (bool, error)
	ListFail2ban() (models.Fail2banListItem, error)
	RemoveFail2ban(ip string) error
	AddFail2ban(ip string) error
	HasFail2banIp(ip string) (bool, error)
}

func (sr *setupRepo) AddressExistsOnServer(addres string) (bool, error) {
	emails, err := sr.ListEmail()
	if err != nil {
		return false, err
	}
	for _, mainEmail := range emails {
		if mainEmail.Address == addres {
			return true, nil
		}
		for _, alias := range mainEmail.Aliases {
			if alias == addres {
				return true, nil
			}
		}
	}
	return false, nil
}

func (sr *setupRepo) AddAlias(aliasAddress, recepientAddres string) error {
	if aliasAddress == "" {
		return errors.New("A alias is required for an email account")
	}

	if recepientAddres == "" {
		return errors.New("A recepient is required for an email account")
	}

	recepientExists, err := sr.AddressExistsOnServer(recepientAddres)
	if err != nil {
		return err
	}
	if !recepientExists {
		return errors.New("Recepient address does not exist on this server!")
	}
	aliasExists, err := sr.AddressExistsOnServer(aliasAddress)
	if err != nil {
		return err
	}
	if aliasExists {
		return errors.New("Alias already in use!")
	}

	parts := strings.Split(sr.commandPrefix+"setup alias add", " ")
	parts = append(parts, aliasAddress)
	parts = append(parts, recepientAddres)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err = proc.Output()
	return err
}
func (sr *setupRepo) GetEmail(addres string) (*models.EmailAccountListItem, error) {
	emails, err := sr.ListEmail()
	if err != nil {
		return nil, err
	}
	for i := range emails {
		if emails[i].Address == addres {
			return emails[i], nil
		}
	}
	return nil, errors.New("Email account not found!")
}
func (sr *setupRepo) HasAlias(aliasAddress, recepientAddres string) (bool, error) {
	emailAcc, err := sr.GetEmail(recepientAddres)
	if err != nil {
		return false, err
	}
	for _, address := range emailAcc.Aliases {
		if address == aliasAddress {
			return true, nil
		}
	}
	return false, nil
}

func (sr *setupRepo) RemoveAlias(aliasAddress, recepientAddres string) error {
	if aliasAddress == "" {
		return errors.New("A alias is required for an email account")
	}

	if recepientAddres == "" {
		return errors.New("A recepient is required for an email account")
	}

	recepientExists, err := sr.AddressExistsOnServer(recepientAddres)
	if err != nil {
		return err
	}
	if !recepientExists {
		return errors.New("Recepient address does not exist on this server!")
	}
	aliasExists, err := sr.AddressExistsOnServer(aliasAddress)
	if err != nil {
		return err
	}
	if !aliasExists {
		return errors.New("Alias does not exist!")
	}

	parts := strings.Split(sr.commandPrefix+"setup alias del", " ")
	parts = append(parts, aliasAddress)
	parts = append(parts, recepientAddres)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err = proc.Output()
	return err
}

func (sr *setupRepo) ListEmail() ([]*models.EmailAccountListItem, error) {
	parts := strings.Split(sr.commandPrefix+"setup email list", " ")
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	t, err := proc.Output()
	if err != nil {
		return nil, err
	}

	emailAddresses := []*models.EmailAccountListItem{}
	for _, line := range strings.Split(string(t), "\n") {
		if len(line) < 3 {
			continue
		}

		if line[0] == '*' {
			lineParts := strings.Split(line[2:], " ( ")
			lineSubParts := strings.Split(lineParts[1], " / ")
			usage := lineSubParts[0]
			lineSubParts2 := strings.Split(lineSubParts[1], " ) [")
			emailAddresses = append(emailAddresses, &models.EmailAccountListItem{
				Address: lineParts[0],
				Aliases: []string{},
				Quota: &models.Quota{
					Usage:           usage,
					Limit:           lineSubParts2[0],
					UsagePercentage: lineSubParts2[1][:len(lineSubParts2[1])-1],
				},
			})
		} else {
			reducedLine := strings.Split(line, "[ aliases -> ")[1]
			reducedLine = reducedLine[:len(reducedLine)-2]
			emailAddresses[len(emailAddresses)-1].Aliases = strings.Split(reducedLine, ", ")
		}
	}

	return emailAddresses, nil
}
func (sr *setupRepo) AddEmail(account models.EmailAccount) error {
	if account.Password == "" {
		return errors.New("A password is required for an email account")
	}

	parts := strings.Split(sr.commandPrefix+"setup email add", " ")
	parts = append(parts, account.Address)
	parts = append(parts, account.Password)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err := proc.Output()
	if err != nil {
		return err
	}

	return nil
}
func (sr *setupRepo) UpdateEmail(account models.EmailAccount) error {
	if account.Password == "" {
		return errors.New("A password is required for an email account")
	}

	parts := strings.Split(sr.commandPrefix+"setup email update", " ")
	parts = append(parts, account.Address)
	parts = append(parts, account.Password)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err := proc.Output()
	if err != nil {
		return err
	}

	return nil
}
func (sr *setupRepo) RemoveEmail(emailAddress string) error {
	parts := strings.Split(sr.commandPrefix+"setup email del", " ")
	parts = append(parts, emailAddress)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err := proc.Output()
	if err != nil {
		return err
	}

	return nil
}
func (sr *setupRepo) ListFail2ban() (models.Fail2banListItem, error) {
	parts := strings.Split(sr.commandPrefix+"setup fail2ban", " ")
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	out := models.Fail2banListItem{}
	t, err := proc.Output()
	if err != nil {
		return out, err
	}

	for _, line := range strings.Split(string(t), "\n") {
		if len(line) < 10 {
			continue
		}
		cleaned := line[10:]
		if cleaned[:7] == "postfix" {
			ips := strings.Split(cleaned[9:], " ")
			out.BannedInPostfix = ips
		}
		if cleaned[:6] == "custom" {
			ips := strings.Split(cleaned[8:], " ")
			out.BannedInCustom = ips
		}
	}

	return out, nil
}
func (sr *setupRepo) AddFail2ban(ip string) error {
	parts := strings.Split(sr.commandPrefix+"setup fail2ban ban", " ")
	parts = append(parts, ip)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err := proc.Output()
	if err != nil {
		return err
	}

	return nil
}
func (sr *setupRepo) RemoveFail2ban(ip string) error {
	parts := strings.Split(sr.commandPrefix+"setup fail2ban unban", " ")
	parts = append(parts, ip)
	proc := exec.Command(parts[0], parts[1:]...)
	proc.Wait()
	_, err := proc.Output()
	if err != nil {
		return err
	}

	return nil
}
func (sr *setupRepo) HasEmail(emailAddress string) (bool, error) {
	emails, err := sr.ListEmail()
	if err != nil {
		return false, err
	}

	for _, email := range emails {
		if email.Address == emailAddress {
			return true, nil
		}
	}

	return false, nil
}
func (sr *setupRepo) HasFail2banIp(ip string) (bool, error) {
	bans, err := sr.ListFail2ban()
	if err != nil {
		return false, err
	}

	for _, currentIp := range bans.BannedInDovecot {
		if currentIp == ip {
			return true, nil
		}
	}

	for _, currentIp := range bans.BannedInPostfix {
		if currentIp == ip {
			return true, nil
		}
	}

	for _, currentIp := range bans.BannedInPostfixSasl {
		if currentIp == ip {
			return true, nil
		}
	}
	for _, currentIp := range bans.BannedInCustom {
		if currentIp == ip {
			return true, nil
		}
	}

	return false, nil
}

func NewSetupRepo(commandPrefix string) SetupRepo {
	return &setupRepo{
		commandPrefix: commandPrefix,
	}
}
