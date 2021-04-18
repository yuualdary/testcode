/* 

$GetData = DB::Table('Transactions')->Select('*',DB::raw('count(Transactions.production_id) as MostProduct' ))
                                    ->Join('Customers','Customers.id','=','Transactions.customer_id')
                                    ->Join('Productions','Productions.id','=','Transactions.production_id')
                                    ->Where('Customers.Gender','Male')
                                    ->groupBy('Transactions.production_id')
                                    ->orderBy('Transactions.production_id','DESC')
                                    ->get();
*/

/*  1 */

SELECT * FROM Transactions
INNER JOIN Customers ON Transactions.customer_id = customers.id 
INNER JOIN productions ON Transactions.production_id = productions.id 
WHERE Customers.Gender = 'Male'  
COUNT  (Transactions.production_id) AS MostProduct From Transactions
GROUP BY Transactions.production_id
ORDER BY Transactions.production_id DESC;



SELECT * FROM Transactions
INNER JOIN Customers ON Transactions.customer_id = customers.id 
INNER JOIN productions ON Transactions.production_id = productions.id 
WHERE Customers.Gender = 'Female'  
COUNT  (Transactions.production_id) AS MostProduct From Transactions
GROUP BY Transactions.production_id
ORDER BY Transactions.production_id DESC;

/* 2 */


SELECT * FROM Transactions
INNER JOIN Customers ON Transactions.customer_id = customers.id 
INNER JOIN productions ON Transactions.production_id = productions.id 
SUM  (productions.Price) AS MostPrice From Transactions
GROUP BY Transactions.customer_id
ORDER BY productions.Price DESC;

/*  */

