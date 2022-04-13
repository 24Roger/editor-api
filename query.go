package main

const getAllProjects string = `
{
	projects(func: has(data)){
    uid
		data
  }
}
`
