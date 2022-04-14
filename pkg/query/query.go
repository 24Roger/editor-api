package query

const GetAllProjects string = `
{
	projects(func: has(data)){
    uid
		data
  }
}
`
