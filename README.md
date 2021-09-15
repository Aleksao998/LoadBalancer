# REGRESION

1. REGISTRATION
---------------------
1.a Registration happy path = YES

1.b Registration fail user already exists = YES


2.LOGIN
---------------------
2.a Login happy path = YES

2.b Login fail email does not exists = YES

2.c Login fail pass not correct = YES

3.JWT
----------------------
3.a Generate token on login and registration = YES

3.b Token validation success = YES

3.c Token invalid = YES

3.d Token expired (3 minutes) = YES

4.LOAD BALANCE
-----------------------
4.a Register worken on load = YES

4.b Try Registrer if fail after minute = YES

4.c De-Rigester worken on shutdown (SIGTERM) = YES

4.d Load Balance no worker = YES

4.e Load Balance happy path (2 workers) = YES

4.f When worker lose connection deregister it = FALSE

4.g Do not allow same address registration multiple time = FALSE

5.BANK ACCOUNT
-------------------------
5.a Create Bank Account happy path = YES

5.b Create fail if user does not exists = YES

5.c Create fail if acc name already exists = YES

5.d Delete Bank acc happy path = YES

5.e No err return if bank acc does not exists od delete = YES

5.f Fetch bank account happy path = YES

5.g Return not exists on fetch if no bank acc = YES

6.EXPENSES
-----------------------
6.a Create Expense happy path = YES

6.b Create fail if expense already exists = YES

6.c Create fail if bank acc not exists = YES

6.d Delete Expense happy path = YES

6.e Return null if no expense to delete = YES

6.f Fetch expense happy path = yes

6.g Return null if no expense 


