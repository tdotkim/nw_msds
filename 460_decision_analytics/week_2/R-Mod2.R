library(lpSolve)
#Problem
#objective function 
objective1<- matrix(c(0.041, 0.008, 0.028, 0.021, 0.012), nrow=1, byrow= TRUE)
#constraints 
constraints_mtrix1<-matrix(c(1,0,0,0,0,
                             0,1,0,0,0,
                             0,0,1,0,0,
                             0,0,0,1,0,
                             0,0,0,0,1,
                             1,1,1,1,1,
                             0.4,0.4,0.4,-0.6,-0.6)
                           , nrow = 7, byrow = TRUE)
constraint_rules1 <- c("<=","<=","<=","<=","<=","<=",">=")
constraint_rhs1<-c(400000,200000,350000,250000,150000,1000000,0)  

#Final solution 0= solution, 2= no solution 
final_solution <- lp(direction = "max", objective.in =objective1, const.mat = constraints_mtrix1,const.dir= constraint_rules1, const.rhs = constraint_rhs1, all.int = FALSE)
print(final_solution$status)
print(final_solution)
final_solution1 <-final_solution$solution
names(final_solution1)<-c("MFR", "EOC", "ER", "SC", "WM")
print(final_solution1)