package cli

import (
    "github.com/fatih/color"
)

func logoPrint() {
    col := color.New(color.FgHiCyan).Add(color.Bold)
	col.Println(`
.......................................................
. . . .  .   . .   . . .. . . .      . .  . . . .   . .
 .  .  . . .  . . . .  .               .  . .  . .  . 
 ::::::::  :::::::::: :::            :::     :::   ::: 
:+:    :+: :+:        :+:          :+: :+:   :+:   :+: 
+:+        +:+        +:+         +:+   +:+   +:+ +:+  
+#++:++#++ +#++:++#   +#+        +#++:++#++:   +#++:   
       +#+ +#+        +#+        +#+     +#+    +#+    
#+#    #+# #+#        #+#        #+#     #+#    #+#    
 ########  ########## ########## ###     ###    ###  
 # # # #   #  #  #  #      #   #  #       #     #  #
  #  #   #  #    # #  #  #   #    #    # #   #   #  
######################################################
	  `)
}