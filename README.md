# multithreading-homework

## UA
# Situation: Distributed hardware
Agenda: В нас є аргокультурний обʼєкт, де рослини підтримуються дронами.
Сенсори на рослинах відправляють дані на сервер, який потім вирішує що робити з ними на основі цих даних.
Кожна рослина має унікальний ID та Data.
Кожен сенсор привʼязаний до своєї рослини та містить її ID, на основі якого відправляє дані до серверу.
Ми маємо три види даних з сенсорів - здоровʼя(стан листя і коріння(0..100%)), гідрацію(0.0..1.0) та кислотність ґрунту.
Для кожної рослини значення низької гідрації та діапазон кислотності індивідуальний.
Поганий стан листя та коріння для всіх однаковий - < 50%.
Також ми можемо спостерегати за станом дронів - заряд(0..100%) та позицію(X, Y).

В нас є база даних, яка представлена у вигляді репозиторієв(PlantRepository, DroneRepository, SensorRepository).
В обробці приймають участь дві структури - Listener та Processor.
Перший чекає повідомлень від сенсорів та відправляє їх через канал до Processor для подальшої обробки.
Processor має логіку обробки та може відправляти команди дронам через DroneRepository.

Mock - директорія, яка містить емуляцію вхідних повідомлень від сенсорів та оновлення стану рослин.

Дрони можуть приймати команди які змінюють стан рослини або команду на заміну рослини з поганим здоровʼям на іншу.

Disclaimer: Дуже рекомендую робити форки - це спрощує процесс перевірки та взагалі стан зробленої роботи.

Завдання: \
А. Реалізуйте метод Listen() у GenericListener, тобто - зчитування повідомлень із сенсорів та відправка їх у зовнішній канал out. \
B. Реалізуйте обробку повідомлень про гідрацію рослин - метод RunProcessor у HydrationProcessor \
C. Реалізуйте обробку повідомлень про кислотність ґрунту рослин - метод RunProcessor у PHProcessor \
D. Реалізуйте обробку повідомлень про стан зродовʼя рослин - метод RunProcessor у HealthProcessor

*E. Якщо почуваєте у себе достатньо сил - реалізуйте вивід стану дронів - їх заряд та позицію.
Дані про це вони повинні пересилати ззовні - реалізуйте також свій mock для цього у PlantsServiceMock.

Якщо в вас є бажання, додавайте свій функціонал або покращте додаток, у окремій гілці.

##ENG

## Situation: Distributed hardware
Agenda: There is an agriculture facility, where plants are maintained by drones.
Server are able to read various info from crops’ sensors and decide, what to do with plants.
Each plant has unique ID and Name.
Each sensor contains plant's ID and Data.
We have three types of data - health(leaves, root status(0...100%)), hydration(0.0...1.0) and ph(0...1000) of plants.
Low hydration and ph is individual to each plant.
Status of root and leaves - bad - < 50%.
Also, we can monitor Charge(0-100%) and Position(X, Y) of drones.

We have database in a view of runtime repos(PlantRepository, DroneRepository, SensorRepository).
Messages are processed through Listener and Processor structs.
Listener waits for messages from sensors and send them to 'out' channel to Processor.
Each Processor process messages by suitable logic and sends commands to drones with DroneRepository.

Mock directory contains emulation of sensors' input and can update plants' state.

Drones can accept commands that changes plant state or plant replace command if health of plant was in bad condition.

Strong recommendation - make tasks in fork repo. With them, it would be easier to check your progress and solution.

Tasks:
A. Your task is to implement method Listen() in GenericListener, that means reading message from sensors and sending them to 'out' channel. \  
B. Your task is to implement hydration messages processing - method RunProcessor in HydrationProcessor. \
C. Your task is to implement ph messages processing - method RunProcessor in PHProcessor. \
D. Your task is to implement health messages processing - method RunProcessor in HealthProcessor.

*E. If you feel powerful enough, implement drones' state output to console, specifically their charge and position.
Data sending can be implemented through mocking in PlantsServiceMock

Feel free to improve program or add additional functionality if you want to, in separate branch.