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
					// oldSkills := contrib.skills

					// requiredLevel := plannedProject.project.rolesList[rolePosition].level
					// if contrib.level

					// for _, role := range plannedProject.project.rolesList {
					// 	role.name
					// }

					contrib.allocated = false
				}
			}
		}

		minNextDay := maxDays
		for _, project := range projects {
			if project.alreadyPlanned || project.bestBefore+project.score-project.nDays < day {
				continue
			}
			plannedProject := &PlannedProject{
				name:         project.name,
				contributors: make([]*Contributor, project.nRoles),
				project:      project,
			}

			for rolePosition, role := range project.rolesList {
				availableContributors := rolesContributor[role.name]

				requiredLevel := role.level
				for _, contributor := range availableContributors {
					if contributor.allocated {
						continue
					}
					contribSkillLevel := contributor.skills[role.name]
					if contribSkillLevel >= requiredLevel {
						//  ||  (contribSkillLevel == requiredLevel-1 && findMenthor(plannedProject.contributors)
						// plannedProject.contributors = append(plannedProject.contributors, contributor)
						plannedProject.contributors[rolePosition] = contributor
						contributor.allocated = true
						break
					}
				}
			}

			// // project roles not filled!!
			if hasunfilledRoles(plannedProject.contributors) {
				for rolePosition, filledRole := range plannedProject.contributors {
					if filledRole != nil {
						continue
					}

					unfilledRole := project.rolesList[rolePosition]

					for _, contributor := range plannedProject.contributors {
						if contributor == nil {
							continue
						}
						contributorSkillLevel := contributor.skills[unfilledRole.name]
						if contributorSkillLevel >= unfilledRole.level {

							availableContributors := rolesContributor[unfilledRole.name]
							for _, availableContributor := range availableContributors {
								if contributor.allocated {
									continue
								}

								if availableContributor.skills[unfilledRole.name] == unfilledRole.level-1 {
									plannedProject.contributors[rolePosition] = contributor
									contributor.allocated = true
									break
								}

							}
						}
						// if filledRole.name == unfilledRole.name {

						// }
					}

					// availableContributors := rolesContributor[unfilledRole.name]
					// for _, contributor := range availableContributors {
					// 	if contributor.allocated {
					// 		continue
					// 	}

					// }
				}
			}
			// if len(plannedProject.contributors) != project.nRoles {

			// 	}
			// 	// for _, role := range project.rolesList {
			// 	// 	requiredLevel := role.level

			// 	// 	availableContributors := rolesContributor[role.name]
			// 	// 	for _, contributor := range availableContributors {
			// 	// 		if contributor.allocated {
			// 	// 			continue
			// 	// 		}
			// 	// 		contribSkillLevel := contributor.skills[role.name]
			// 	// 		if contribSkillLevel == requiredLevel-1 {
			// 	// 			existsMenthor()

			// 	// 			//  ||  (contribSkillLevel == requiredLevel-1 && findMenthor(plannedProject.contributors)
			// 	// 			plannedProject.contributors = append(plannedProject.contributors, contributor)
			// 	// 			contributor.allocated = true
			// 	// 			break
			// 	// 		}
			// 	// 	}
			// 	}
			// }

			// can't use the project!
			// if len(plannedProject.contributors) != project.nRoles {
			if hasunfilledRoles(plannedProject.contributors) {
				for _, contrib := range plannedProject.contributors {
					if contrib == nil {
						continue
					}
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

func hasunfilledRoles(contributors []*Contributor) bool {
	for _, contrib := range contributors {
		if contrib == nil {
			return true
		}
	}
	return false
}
