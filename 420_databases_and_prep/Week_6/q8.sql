select 
	t.tm_month as month,
	count(s.sale_units) as productsalescount,
	sum(s.sale_units*s.sale_price) as totalsales
from dwdaysalesfact as s
	inner join dwtime as t on s.tm_id=t.tm_id
group by month
order by month