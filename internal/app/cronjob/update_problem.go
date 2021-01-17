package cronjob

import (
	"context"
)

const (
	SelectRandOne = `create temporary table unpicked (select * from problems where id not in (select problem_id from user_problems where user_id = 1));
		select count(*) into @C from unpicked;
		set @Y = floor(@C * rand());
		set @sql = concat("select * from unpicked limit ", @Y, ",1");
		prepare stmt from @sql;
		execute stmt;
		DEALLOCATE prepare stmt;`
)

func UpdateProblem(ctx context.Context) {

}
