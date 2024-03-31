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

	ListFail2ban() (models.Fail2banListItem, error)
	RemoveFail2ban(ip string) error
	AddFail2ban(ip string) error
	HasFail2banIp(ip string) (bool, error)
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
		cleaned := line[11:]
		if cleaned[:7] == "postfix" {
			ips := strings.Split(cleaned[9:], " ")
			out.BannedInPostfix = ips
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

	return false, nil
}

func NewSetupRepo(commandPrefix string) SetupRepo {
	return &setupRepo{
		commandPrefix: commandPrefix,
	}
}
