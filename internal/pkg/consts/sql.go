package consts

const (
	GetDailyProblemSQL = `select * from problems where id in 
		(select problem_id from user_problems 
		where user_id = ? and picked = true and finished = false)
		order by problem_type`

	SelectUserProblemTypeSQL = `select distinct(problem_type) from user_problems
		where user_id = ? and (picked = false or finished = true)`

	CountUnPickedProblemSQL = `select count(*) from user_problems 
		where user_id = ? and problem_type = ? and picked = false`

	SelectRandUnPickedProblemSQL = `select * from 
		(select * from user_problems 
		where user_id = ? and problem_type = ? and picked = false) as unpicked 
		limit ?,1`

	CountFinishedProblemSQL = `select count(*) from user_problems 
		where user_id = ? and problem_type = ? and picked = true and finished = true`

	SelectRandFinishedProblemSQL = `select * from 
		(select * from user_problems 
		where user_id = ? and problem_type = ? and picked = true and finished = true) as finished 
		limit ?,1`
)
