package main

type PlannedProject struct {
	name         string
	contributors []string
	project      *Project
}

func algorithm(
	config *Config,
	contributors []*Contributor,
	projects []*Project,
	rolesContributor map[string][]*Contributor,
) []*PlannedProject {
	plannedProjects := make([]*PlannedProject, 0)

	for _, project := range projects {
		plannedProject := &PlannedProject{
			name:         project.name,
			contributors: make([]string, 0),
		}

		for _, role := range project.rolesList {
			availableContributors := rolesContributor[role.name]

			requiredLevel := role.level
			for _, contributor := range availableContributors {
				contribSkillLevel := contributor.skills[role.name]
				if contribSkillLevel >= requiredLevel {
					plannedProject.contributors = append(plannedProject.contributors, contributor.name)
					break
				}
			}

		}

		// Non abbiamo riempito i ruoli!
		if len(plannedProject.contributors) != project.nRoles {
			break
		}

		// append
		plannedProjects = append(plannedProjects, plannedProject)
	}
	return plannedProjects
}
