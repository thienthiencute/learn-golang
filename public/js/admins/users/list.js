var user_dt;

const userList = function(){
    const initial = function (){
        new DataTable('#user_table');
        

    };
    return {
        init: function(){
            initial();

        },
    };
};

document.addEventListener("DOMContentLoaded", function (event){
    userList().init();
});