package main

import "sort"

type PlannedProject struct {
	name         string
	startDay     int
	ended        bool
	contributors []*Contributor
	project      *Project
}

func algorithm(
	maxDays int,
	config *Config,
	contributors []*Contributor,
	projects []*Project,
	rolesContributor map[string][]*Contributor,
) []*PlannedProject {
	plannedProjects := make([]*PlannedProject, 0)

	sort.Slice(projects, func(a, b int) bool {
		projectA := projects[a]
		projectB := projects[b]

		scoreA := projectA.bestBefore
		scoreB := projectB.bestBefore

		return scoreA < scoreB
	})

	nextDay := 0
	for day := 0; day < maxDays || len(plannedProjects) >= len(projects); day++ {
		if day < nextDay {
			continue
		}
		for _, plannedProject := range plannedProjects {
			if plannedProject.ended {
				continue
			}

			if day >= plannedProject.startDay+plannedProject.project.nDays {
				plannedProject.ended = true
				for _, contrib := range plannedProject.contributors {
					contrib.allocated = false
				}
			}
		}

		// for _, project := range projects {
		// 	fmt.Printf("> Project %+v\n", project)
		// 	for _, role := range project.rolesMap {
		// 		fmt.Printf(">> Role %+v\n", role)
		// 	}
		// }

		minNextDay := maxDays
		for _, project := range projects {
			// fmt.Printf("%d - PROCESSING PROJECT %s\n", day, project.name)
			if project.alreadyPlanned || project.bestBefore+project.score-project.nDays < day {
				continue
			}
			plannedProject := &PlannedProject{
				name:         project.name,
				contributors: make([]*Contributor, 0),
				project:      project,
			}

			for _, role := range project.rolesList {
				availableContributors := rolesContributor[role.name]

				requiredLevel := role.level
				for _, contributor := range availableContributors {
					if contributor.allocated {
						continue
					}
					contribSkillLevel := contributor.skills[role.name]
					if contribSkillLevel >= requiredLevel {
						plannedProject.contributors = append(plannedProject.contributors, contributor)
						contributor.allocated = true
						break
					}
				}
			}

			// fmt.Printf("SHOULD BE PLANNING PROJECT %s %+v %+v\n", project.name, plannedProject, plannedProject.contributors)
			// Non abbiamo riempito i ruoli!
			if len(plannedProject.contributors) != project.nRoles {
				for _, contrib := range plannedProject.contributors {
					contrib.allocated = false
				}
				continue
			}

			// append
			plannedProject.startDay = day

			plannedProject.project.alreadyPlanned = true
			plannedProjects = append(plannedProjects, plannedProject)
		}

		for _, plannedProject := range plannedProjects {
			if day+plannedProject.project.nDays < minNextDay {
				minNextDay = day + plannedProject.project.nDays
			}
		}

		nextDay = minNextDay
	}

	// for _, project := range projects {
	// 	plannedProject := &PlannedProject{
	// 		name:         project.name,
	// 		contributors: make([]string, 0),
	// 	}

	// 	for _, role := range project.rolesList {
	// 		availableContributors := rolesContributor[role.name]

	// 		requiredLevel := role.level
	// 		for _, contributor := range availableContributors {
	// 			contribSkillLevel := contributor.skills[role.name]
	// 			if contribSkillLevel >= requiredLevel {
	// 				plannedProject.contributors = append(plannedProject.contributors, contributor.name)
	// 				break
	// 			}
	// 		}
	// 	}

	// 	// Non abbiamo riempito i ruoli!
	// 	if len(plannedProject.contributors) != project.nRoles {
	// 		break
	// 	}

	// 	// append
	// 	plannedProjects = append(plannedProjects, plannedProject)
	// }
	return plannedProjects
}
