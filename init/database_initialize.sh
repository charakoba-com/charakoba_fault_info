host='localhost'
user='root'
pass=''
while getopts h:u:p: OPT; do
    case $OPT in
        h|+h)
            host="$OPTARG"
            ;;
        u|+u)
            user="$OPTARG"
            ;;
        p|+p)
            pass="$OPTARG"
            ;;
        *)
            echo "usage: `basename $0` [+-u ARG] [+-p ARG} [--] ARGS..."
            exit 2
    esac
done
shift `expr $OPTIND - 1`
OPTIND=1

create_database="
CREATE DATABASE fault_info;
"

create_table="
CREATE TABLE fault_info.fault_info_log (
       id INT PRIMARY KEY AUTO_INCREMENT,
       type ENUM('mainenance', 'event') NOT NULL,
       service VARCHAR(64) NOT NULL,
       begin DATETIME NOT NULL,
       end DATETIME,
       detail VARCHAR(512) NOT NULL DEFAULT ''
);
"

mysql -h ${host} -u ${user} --password=${pass} -e "${create_database}"
mysql -h ${host} -u ${user} --password=${pass} -e "${create_table}"
