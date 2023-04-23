select 
	p.p_category as category,
	p.p_code as code,
	p.p_descript as description,
	sum(s.sale_units) as unitcount
from dwdaysalesfact as s
	inner join dwtime as t on s.tm_id=t.tm_id
	inner join dwproduct as p on s.p_code=p.p_code
group by p.p_category,p.p_code,p.p_descript
order by unitcount desc