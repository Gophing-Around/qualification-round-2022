package main

import "fmt"

type Config struct {
	contributors int
	projects     int
}

type Contributor struct {
	name      string
	nSkills   int
	allocated bool
	skills    map[string]int
}

type Project struct {
	name       string
	nDays      int
	score      int
	bestBefore int
	nRoles     int
	rolesList  []*Role
	rolesMap   map[string]*Role

	alreadyPlanned bool
}

type Role struct {
	id       int
	name     string
	level    int
	mentored bool
}

func buildInput(inputSet string) (int, *Config, []*Contributor, []*Project, map[string][]*Contributor) {
	lines := splitNewLines(inputSet)
	configLine := splitSpaces(lines[0])
	fmt.Printf("Config line: %v\n", configLine)

	config := &Config{
		contributors: toint(configLine[0]),
		projects:     toint(configLine[1]),
	}

	contributors := make([]*Contributor, 0)
	projects := make([]*Project, 0)
	rolesContributor := make(map[string][]*Contributor)

	i := 1

	maxDays := 0

	k := 0
	for ; k < config.contributors; k++ {
		contribLine := splitSpaces(lines[i])
		contrib := &Contributor{
			name:    contribLine[0],
			nSkills: toint(contribLine[1]),
		}

		i++
		skill := make(map[string]int)
		j := 0
		for ; j < contrib.nSkills; j++ {
			skillLine := splitSpaces(lines[i+j])
			skillName := skillLine[0]
			skill[skillName] = toint(skillLine[1])

			list, ok := rolesContributor[skillName]
			if !ok {
				list = make([]*Contributor, 0)
			}
			list = append(list, contrib)
			rolesContributor[skillName] = list
		}
		i += j
		contrib.skills = skill
		contributors = append(contributors, contrib)
	}

	for k := 0; k < config.projects; k++ {
		projectLine := splitSpaces(lines[i])
		i++

		project := Project{
			name:       projectLine[0],
			nDays:      toint(projectLine[1]),
			score:      toint(projectLine[2]),
			bestBefore: toint(projectLine[3]),
			nRoles:     toint(projectLine[4]),
		}

		if maxDays < project.bestBefore+project.score-project.nDays {
			maxDays = project.bestBefore + project.score - project.nDays
		}

		j := 0
		roles := make(map[string]*Role)
		rolesList := make([]*Role, 0)
		for ; j < project.nRoles; j++ {
			roleLine := splitSpaces(lines[i+j])
			role := Role{
				id:    j,
				name:  roleLine[0],
				level: toint(roleLine[1]),
			}

			roles[roleLine[0]] = &role
			rolesList = append(rolesList, &role)
		}
		i += j
		project.rolesMap = roles
		project.rolesList = rolesList
		projects = append(projects, &project)
	}

	return maxDays, config, contributors, projects, rolesContributor
}

func buildOutput(plannedProjects []*PlannedProject) string {
	result := fmt.Sprintf("%d\n", len(plannedProjects))
	for _, project := range plannedProjects {
		result += fmt.Sprintf("%s\n", project.name)
		for _, contrib := range project.contributors {
			result += fmt.Sprintf("%s ", contrib.name)
		}
		result += "\n"
	}
	return result
}
