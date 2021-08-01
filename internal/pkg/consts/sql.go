package consts

const (
	// 查询某人今天未完成的任务
	GetDailyProblemSQL = `select * from problems where id in 
		(select problem_id from user_problems 
		where user_id = ? and picked = true and finished = false and deleted_at is null)
		order by type`

	SelectUserProblemTypeSQL = `select distinct(problem_type) from user_problems
		where user_id = ? and (picked = false or finished = true) and deleted_at is null`

	CountUnPickedProblemSQL = `select count(*) from user_problems 
		where user_id = ? and problem_type = ? and picked = false and deleted_at is null`

	SelectRandUnPickedProblemSQL = `select * from 
		(select * from user_problems 
		where user_id = ? and problem_type = ? and picked = false and deleted_at is null) as unpicked 
		limit ?,1`

	CountFinishedProblemSQL = `select count(*) from user_problems 
		where user_id = ? and problem_type = ? and picked = true and finished = true and deleted_at is null`

	SelectRandFinishedProblemSQL = `select * from 
		(select * from user_problems 
		where user_id = ? and problem_type = ? and picked = true and finished = true and deleted_at is null) as finished 
		limit ?,1`

	SelectProblemUnplannedSQL = `select * from problems where 
		(creator_id = ? or is_public = true)
		and id not in (select problem_id from user_problems where user_id = ? and deleted_at is null)
		order by created_at desc`

	SelectProblemPlannedSQL = `select * from problems where id in
		(select problem_id from user_problems where user_id = ? and deleted_at is null)
		order by created_at desc`

	// 查询某人迄今已做多少道题、已做多少次题
	SelectDoneProblemSQL = `select count(times) num, coalesce(sum(times), 0) times from user_problems
		where user_id = ? and times > 0`

	// 查询某人今天有多少题没做、已做多少题
	SelectTodayWorkloadSQL = `select count(1) cnt from user_problems where 
		user_id = ? and picked = true and finished = false 
		union all
		select count(1) cnt from user_problems where 
		user_id = ? and picked = true and finished = true and to_days(now()) = to_days(updated_at)`

	//`select problems.*, t.picked, t.pick_time, t.finished, t.times. from (select * from user_problems where user_id = ? and picked = true and finished = false) as t left join problems on t.problem_id = problems.id order by problem_type`
)
