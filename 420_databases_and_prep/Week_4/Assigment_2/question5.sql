-- SQLite
SELECT 
    CUS_CODE,
    CUS_FNAME || ' ' || CUS_LNAME AS CUS_NAME,
    CUS_BALANCE 
FROM CUSTOMER
WHERE CUS_BALANCE > 500

