# partone
setting up directory and serving static files using go-chi

# directory layout and concepts

I come from a Asp.net MVC background.  as such there is some conventions:

* **controllers**: this is where URl routes are handled.  I use Root to refer to paths at /.  
Otherwise, sub paths are handled by named controllers.  
/admin is handled by admin.go  
/admin/DeleteSomething would be a route defined and handled in the admin.go file.   
* **views**: views are laid out to match the controllers.  a view such as views/admin/ListUsers.tmpl would be found used in the admin controller, preferrably in AdminListUsersHandler()
* **public**: list of public static files, such as css, images, javascript.

In later parts I will be using this same layout, and extending it as the application becomes more complex.
