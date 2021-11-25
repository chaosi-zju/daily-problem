package consts

const (
	// 查询某人今天未完成的任务
	GetDailyProblemSQL = `select * from problems where id in 
		(select problem_id from user_problems 
		where user_id = ? and picked = true and finished = false and deleted_at is null)
		order by type desc, sub_type, name asc`

	GetCommonDailyProblemSQL = `select * from problems where id in 
		(select problem_id from user_problems where user_id = 1 and picked = true and finished = false and deleted_at is null
		 union
		 select problem_id from user_problem_logs where user_id = 1 and action = 'finish' and to_days(date_sub(now(), interval ? hour)) = to_days(action_time))
		 order by type desc`

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
	SelectDoneProblemSQL = `SELECT
		    sum(IF(problem_type = 'algorithm' and times > 0, 1, 0)) algorithm_done_num,
			sum(IF(problem_type = 'algorithm' and (deleted_at is null or times > 0), 1, 0)) algorithm_all_num,
			sum(IF(problem_type = 'algorithm' and times > 0, times, 0)) algorithm_done_times,
			sum(IF(problem_type != 'algorithm' and times > 0, 1, 0)) other_done_num,
			sum(IF(problem_type != 'algorithm' and (deleted_at is null or times > 0), 1, 0)) other_all_num,
			sum(IF(problem_type != 'algorithm' and times > 0, times, 0)) other_done_times
		FROM
			user_problems 
		WHERE
			user_id = ?`

	// 查询某人今天有多少题没做、已做多少题
	SelectTodayWorkloadSQL = `select count(1) cnt from user_problems where 
		user_id = ? and picked = true and finished = false and deleted_at is null
		union all
		select count(1) cnt from user_problem_logs where 
		user_id = ? and action = 'finish' and to_days(date_sub(now(), interval ? hour)) = to_days(action_time)`

	// 查询各个用户每天分别做了几道哪种类型的题 limit 1000
	SelectFinishInfoSQL = `SELECT date, name as user, GROUP_CONCAT(CONCAT(count, '道', problem_type, '题') ORDER BY problem_type SEPARATOR '，') amount FROM
	(
		SELECT date, user_id, name, problem_type, count(1) AS count 
		FROM (
				select date_format(date_sub(a.action_time, INTERVAL ? HOUR), '%Y-%m-%d') AS date, a.*, b.name 
				from user_problem_logs a left join users b on a.user_id = b.id
			) c
		GROUP BY date, user_id, problem_type ORDER BY date DESC, user_id, problem_type LIMIT 1000 
	) d 
	GROUP BY date, user_id ORDER BY date DESC, user_id`

	SelectFinishChartSQL = `SELECT date, problem_type, count(1) count FROM
	(
		SELECT date_format(date_sub(action_time, INTERVAL ? HOUR), '%m-%d') AS date, a.*
		FROM user_problem_logs a WHERE user_id = ? AND action = 'finish'
	) b 
	GROUP BY date, problem_type ORDER BY date, problem_type`

	SelectAllProblems = `select * from problems where creator_id = ?`

	//`select problems.*, t.picked, t.pick_time, t.finished, t.times. from (select * from user_problems where user_id = ? and picked = true and finished = false) as t left join problems on t.problem_id = problems.id order by problem_type`
)
