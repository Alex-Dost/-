- ls - Список содержимого каталога. используется для вывода списка файлов и каталогов
Команда ls -l | more – помогает разбить вывод на страницы, чтобы вы могли просматривать страницу за страницей
- cd /var/log – Изменить текущий каталог
Команда ls -I, позволяет увидеть более подробную информацию о содержимом каталога.
- grep - Найти текст в файле. Команда grep просматривает все файлы, чтобы найти фрагмент текста, который вы ищете.
- su / sudo - Команда su изменяет оболочку, которая будет использоваться в качестве суперпользователя, и пока вы не используете команду exit, вы можете оставаться суперпользователем. Пример - команда shutdown безопасно отключает компьютер.
- pwd - текущий каталог - Один из способов определить каталог в котором вы работаете является команда pwdpwd
- passwd - используется для изменения пароля учетной записи пользователя. passwd [имя пользователя] - изменяет пароль для пользователя.
- mv - переместить файл. Чтобы переместить или переименовать файл, вы должны использовать команду mv . Ниже имя файла меняется с first.txt на second.txt.
- cp - скопировать файл. В случае , если нужна копия файла second.txt в том же каталоге , вы должны использовать команду ср. можно использовать ls - l, чтобы увидеть созданный новый файл. Два файла будут точно одинакового размера
- rм - Эта команда используется для удаления файлов в каталоге. Каталог не может быть удален, если он не пуст. rm [имя файла] . rm –r удаляет все содержимое каталога, а также и сам каталог.
- mkdir - создать каталог. mkdir [имя каталога], если вы хотите создать каталог с именем «myproject» . mkdir myproject
- chmod - Изменить права доступа для каталога. Файлы могут иметь права r - чтение, w - запись и x - выполнение.
Например:
CHMOD mode FILE
chmod 744 script.sh
Первый номер обозначает пользователя, который связан с файлом
Второе число для группы, связанной с файлом
Третий номер связан со всеми, кто не является частью пользователя или группы
- chown – используется для изменения владельца файла / папки или даже нескольких файлов / папок для указанного пользователя / группы.
chown owner_name имя_файла
Предположим, что если вы пользователь с именем user1 и хотите сменить владельца на root, используйте «sudo» перед синтаксисом.
$ sudo chown root script.sh
- cat - (сокращение от «concatenate») является одной из наиболее часто используемых команд в Linux.
cat позволяет создавать один или несколько файлов, просматривать содержимое файла, объединяет файлы и перенаправляет вывод в терминал или файлы.
Вывод покажет все содержимое файла (ов).
- echo – используется для отображения текста или строки для стандартного вывода или файла.
$ echo «Это статья об основных командах Linux»
Опция echo –e действует как интерпретация escape-символов, которые имеют обратную косую черту.
- wc - (word count) в операционной системе Linux используется для определения количества новых строк, слов, количества байтов и символов в файле.
wc [опции] имена файлов
wc -l : печатает количество строк в файле.
wc -w : печатает количество слов в файле.
wc -c : отображает количество байтов в файле.
wc -m : печатает количество символов из файла.
wc -L : печатает только длину самой длинной строки в файле.
- man – используется для просмотра страниц справочного руководства для команд / программ.
- history – используется для отображения ранее использованных команд или для получения информации о командах, выполняемых пользователем.
- clear – позволяет очистить экран терминала.
- apt –get - это мощный и бесплатный менеджер пакетов для систем Debian / Ubuntu.
используется для установки новых пакетов программного обеспечения, удаления доступных пакетов программного обеспечения, обновления существующих пакетов программного обеспечения, а также обновления всей операционной системы. apt - обозначает усовершенствованный упаковочный инструмент.  
- reboot – может использоваться для остановки, выключения или перезагрузки системы.