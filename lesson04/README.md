# ДЗ 04

## 1. С помощью пула воркеров написать программу,которая запускает 1000 горутин, каждая из которых увеличивает число на 1. Дождаться завершения всех горутин и убедиться, что при каждом запуске программы итоговое число равно 1000
### 1.1 Некорректное решение (параллельная работа с общей переменной) 
Создаем буферизированный канал на 100 значений, для ограничения количества одновременно работающих воркеров:  
``` golang
var workers = make(chan struct{}, 100)
```

Объявляем переменную счетчика:  
``` golang
var counter int
```

Запускаем 1000 горутин в цикле, при этом одновременно могут работать только 100 воркеров.
```golang
for i := 1; i <= 1000; i++ {
    workers <- struct{}{}
    go func(job int) {
        defer func() {
            <-workers // Read from channel to unlock another waiting goroutines
        }()
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("Job number: %d; Counter: %d, Number of run goroutines %d\n", job, counter, len(workers))
        counter++
    }(i)
}
```

Остальные ожидают освобождение канала (defer после окончания работы горутины).  
Установил 0,5 секунды для симуляции длительного выполнения функции и визуализации блокировок.

<details>
<summary>Итог выполнения:  </summary>
<pre>
Job number: 15; Counter: 0, Number of run goroutines 100 
Job number: 8; Counter: 0, Number of run goroutines 100  
Job number: 22; Counter: 0, Number of run goroutines 100 
Job number: 2; Counter: 0, Number of run goroutines 100  
Job number: 24; Counter: 0, Number of run goroutines 100 
Job number: 30; Counter: 0, Number of run goroutines 100 
Job number: 29; Counter: 0, Number of run goroutines 100 
Job number: 27; Counter: 0, Number of run goroutines 100 
Job number: 26; Counter: 0, Number of run goroutines 100 
Job number: 28; Counter: 0, Number of run goroutines 100 
Job number: 21; Counter: 0, Number of run goroutines 100 
Job number: 23; Counter: 0, Number of run goroutines 100 
Job number: 17; Counter: 0, Number of run goroutines 100 
Job number: 16; Counter: 0, Number of run goroutines 100 
Job number: 11; Counter: 0, Number of run goroutines 100
Job number: 14; Counter: 0, Number of run goroutines 100
Job number: 10; Counter: 0, Number of run goroutines 100
Job number: 9; Counter: 0, Number of run goroutines 100
Job number: 32; Counter: 0, Number of run goroutines 100
Job number: 100; Counter: 0, Number of run goroutines 100
Job number: 99; Counter: 0, Number of run goroutines 100
Job number: 98; Counter: 0, Number of run goroutines 100
Job number: 97; Counter: 0, Number of run goroutines 100
Job number: 96; Counter: 0, Number of run goroutines 100
Job number: 95; Counter: 0, Number of run goroutines 100
Job number: 94; Counter: 0, Number of run goroutines 100
Job number: 93; Counter: 0, Number of run goroutines 100
Job number: 92; Counter: 0, Number of run goroutines 100
Job number: 91; Counter: 0, Number of run goroutines 100
Job number: 90; Counter: 0, Number of run goroutines 100
Job number: 89; Counter: 0, Number of run goroutines 100
Job number: 88; Counter: 0, Number of run goroutines 100
Job number: 87; Counter: 0, Number of run goroutines 100
Job number: 86; Counter: 0, Number of run goroutines 100
Job number: 85; Counter: 0, Number of run goroutines 100
Job number: 84; Counter: 0, Number of run goroutines 100
Job number: 83; Counter: 0, Number of run goroutines 100
Job number: 82; Counter: 0, Number of run goroutines 100
Job number: 81; Counter: 0, Number of run goroutines 100
Job number: 80; Counter: 0, Number of run goroutines 100
Job number: 79; Counter: 0, Number of run goroutines 100
Job number: 55; Counter: 0, Number of run goroutines 100
Job number: 54; Counter: 0, Number of run goroutines 100
Job number: 53; Counter: 0, Number of run goroutines 100
Job number: 52; Counter: 0, Number of run goroutines 100
Job number: 51; Counter: 0, Number of run goroutines 100
Job number: 50; Counter: 0, Number of run goroutines 100
Job number: 49; Counter: 0, Number of run goroutines 100
Job number: 56; Counter: 0, Number of run goroutines 100
Job number: 48; Counter: 0, Number of run goroutines 100
Job number: 47; Counter: 0, Number of run goroutines 100
Job number: 78; Counter: 0, Number of run goroutines 100
Job number: 46; Counter: 0, Number of run goroutines 100
Job number: 77; Counter: 0, Number of run goroutines 100
Job number: 45; Counter: 0, Number of run goroutines 100
Job number: 44; Counter: 0, Number of run goroutines 100
Job number: 76; Counter: 0, Number of run goroutines 100
Job number: 43; Counter: 0, Number of run goroutines 100
Job number: 75; Counter: 0, Number of run goroutines 100
Job number: 42; Counter: 0, Number of run goroutines 100
Job number: 74; Counter: 0, Number of run goroutines 100
Job number: 41; Counter: 0, Number of run goroutines 100
Job number: 73; Counter: 0, Number of run goroutines 100
Job number: 40; Counter: 0, Number of run goroutines 100
Job number: 39; Counter: 0, Number of run goroutines 100
Job number: 38; Counter: 0, Number of run goroutines 100
Job number: 72; Counter: 0, Number of run goroutines 100
Job number: 37; Counter: 0, Number of run goroutines 100
Job number: 71; Counter: 0, Number of run goroutines 100
Job number: 36; Counter: 0, Number of run goroutines 100
Job number: 70; Counter: 0, Number of run goroutines 100
Job number: 35; Counter: 0, Number of run goroutines 100
Job number: 69; Counter: 0, Number of run goroutines 100
Job number: 34; Counter: 0, Number of run goroutines 100
Job number: 68; Counter: 0, Number of run goroutines 100
Job number: 33; Counter: 0, Number of run goroutines 100
Job number: 61; Counter: 0, Number of run goroutines 100
Job number: 65; Counter: 0, Number of run goroutines 100
Job number: 60; Counter: 0, Number of run goroutines 100
Job number: 67; Counter: 0, Number of run goroutines 100
Job number: 59; Counter: 0, Number of run goroutines 100
Job number: 66; Counter: 0, Number of run goroutines 100
Job number: 58; Counter: 0, Number of run goroutines 100
Job number: 57; Counter: 0, Number of run goroutines 100
Job number: 64; Counter: 0, Number of run goroutines 100
Job number: 63; Counter: 0, Number of run goroutines 100
Job number: 1; Counter: 0, Number of run goroutines 100
Job number: 6; Counter: 0, Number of run goroutines 100
Job number: 31; Counter: 0, Number of run goroutines 100
Job number: 5; Counter: 0, Number of run goroutines 100
Job number: 4; Counter: 0, Number of run goroutines 100
Job number: 3; Counter: 0, Number of run goroutines 100
Job number: 19; Counter: 0, Number of run goroutines 100
Job number: 12; Counter: 0, Number of run goroutines 100
Job number: 13; Counter: 0, Number of run goroutines 100
Job number: 25; Counter: 0, Number of run goroutines 100
Job number: 7; Counter: 0, Number of run goroutines 100
Job number: 20; Counter: 0, Number of run goroutines 100 
Job number: 18; Counter: 0, Number of run goroutines 100
Job number: 62; Counter: 0, Number of run goroutines 100
Job number: 133; Counter: 100, Number of run goroutines 100 
Job number: 138; Counter: 100, Number of run goroutines 100 
Job number: 153; Counter: 102, Number of run goroutines 100
Job number: 127; Counter: 100, Number of run goroutines 100
Job number: 136; Counter: 100, Number of run goroutines 100
Job number: 137; Counter: 100, Number of run goroutines 100
Job number: 116; Counter: 100, Number of run goroutines 100
Job number: 154; Counter: 107, Number of run goroutines 100
Job number: 107; Counter: 100, Number of run goroutines 100
Job number: 108; Counter: 100, Number of run goroutines 100
Job number: 109; Counter: 100, Number of run goroutines 100
Job number: 111; Counter: 100, Number of run goroutines 100
Job number: 112; Counter: 100, Number of run goroutines 100
Job number: 113; Counter: 100, Number of run goroutines 100
Job number: 110; Counter: 100, Number of run goroutines 100
Job number: 114; Counter: 100, Number of run goroutines 100
Job number: 115; Counter: 100, Number of run goroutines 100
Job number: 140; Counter: 100, Number of run goroutines 100
Job number: 152; Counter: 100, Number of run goroutines 100
Job number: 163; Counter: 119, Number of run goroutines 100
Job number: 141; Counter: 100, Number of run goroutines 100
Job number: 102; Counter: 100, Number of run goroutines 100
Job number: 103; Counter: 100, Number of run goroutines 100
Job number: 165; Counter: 123, Number of run goroutines 100
Job number: 143; Counter: 100, Number of run goroutines 100
Job number: 104; Counter: 100, Number of run goroutines 100
Job number: 126; Counter: 100, Number of run goroutines 100
Job number: 168; Counter: 127, Number of run goroutines 100
Job number: 132; Counter: 100, Number of run goroutines 100
Job number: 139; Counter: 100, Number of run goroutines 100
Job number: 128; Counter: 100, Number of run goroutines 100
Job number: 130; Counter: 100, Number of run goroutines 100
Job number: 129; Counter: 100, Number of run goroutines 100
Job number: 122; Counter: 100, Number of run goroutines 100
Job number: 148; Counter: 100, Number of run goroutines 100
Job number: 174; Counter: 135, Number of run goroutines 100
Job number: 145; Counter: 100, Number of run goroutines 100
Job number: 118; Counter: 100, Number of run goroutines 100
Job number: 146; Counter: 100, Number of run goroutines 100
Job number: 176; Counter: 139, Number of run goroutines 100
Job number: 119; Counter: 100, Number of run goroutines 100
Job number: 123; Counter: 100, Number of run goroutines 100
Job number: 182; Counter: 142, Number of run goroutines 100
Job number: 124; Counter: 100, Number of run goroutines 100
Job number: 121; Counter: 100, Number of run goroutines 100
Job number: 125; Counter: 100, Number of run goroutines 100
Job number: 134; Counter: 100, Number of run goroutines 100
Job number: 150; Counter: 100, Number of run goroutines 100
Job number: 149; Counter: 100, Number of run goroutines 100
Job number: 186; Counter: 149, Number of run goroutines 100
Job number: 144; Counter: 100, Number of run goroutines 100
Job number: 151; Counter: 100, Number of run goroutines 100
Job number: 106; Counter: 100, Number of run goroutines 100
Job number: 105; Counter: 100, Number of run goroutines 100
Job number: 155; Counter: 107, Number of run goroutines 100
Job number: 156; Counter: 107, Number of run goroutines 100
Job number: 157; Counter: 110, Number of run goroutines 100
Job number: 158; Counter: 110, Number of run goroutines 100
Job number: 159; Counter: 110, Number of run goroutines 100
Job number: 160; Counter: 114, Number of run goroutines 100
Job number: 161; Counter: 114, Number of run goroutines 100
Job number: 101; Counter: 100, Number of run goroutines 100
Job number: 164; Counter: 119, Number of run goroutines 100
Job number: 162; Counter: 119, Number of run goroutines 100
Job number: 142; Counter: 100, Number of run goroutines 100
Job number: 166; Counter: 123, Number of run goroutines 100
Job number: 131; Counter: 100, Number of run goroutines 100
Job number: 167; Counter: 127, Number of run goroutines 100
Job number: 171; Counter: 130, Number of run goroutines 100
Job number: 172; Counter: 130, Number of run goroutines 100
Job number: 170; Counter: 130, Number of run goroutines 100
Job number: 169; Counter: 130, Number of run goroutines 100
Job number: 117; Counter: 100, Number of run goroutines 100
Job number: 173; Counter: 135, Number of run goroutines 100
Job number: 175; Counter: 135, Number of run goroutines 100
Job number: 147; Counter: 100, Number of run goroutines 100
Job number: 178; Counter: 139, Number of run goroutines 100
Job number: 177; Counter: 139, Number of run goroutines 100
Job number: 120; Counter: 100, Number of run goroutines 100
Job number: 180; Counter: 142, Number of run goroutines 100
Job number: 181; Counter: 142, Number of run goroutines 100
Job number: 179; Counter: 142, Number of run goroutines 100
Job number: 183; Counter: 145, Number of run goroutines 100
Job number: 184; Counter: 145, Number of run goroutines 100
Job number: 185; Counter: 145, Number of run goroutines 100
Job number: 189; Counter: 149, Number of run goroutines 100
Job number: 135; Counter: 100, Number of run goroutines 100
Job number: 187; Counter: 149, Number of run goroutines 100
Job number: 188; Counter: 149, Number of run goroutines 100
Job number: 190; Counter: 152, Number of run goroutines 100
Job number: 191; Counter: 152, Number of run goroutines 100
Job number: 192; Counter: 152, Number of run goroutines 100
Job number: 193; Counter: 152, Number of run goroutines 100
Job number: 195; Counter: 156, Number of run goroutines 100
Job number: 196; Counter: 156, Number of run goroutines 100
Job number: 194; Counter: 156, Number of run goroutines 100
Job number: 198; Counter: 160, Number of run goroutines 100
Job number: 197; Counter: 160, Number of run goroutines 100
Job number: 200; Counter: 165, Number of run goroutines 100
Job number: 199; Counter: 165, Number of run goroutines 100
Job number: 250; Counter: 200, Number of run goroutines 100 
Job number: 299; Counter: 200, Number of run goroutines 100 
Job number: 201; Counter: 200, Number of run goroutines 100
Job number: 214; Counter: 200, Number of run goroutines 100
Job number: 217; Counter: 200, Number of run goroutines 100
Job number: 215; Counter: 200, Number of run goroutines 100
Job number: 216; Counter: 200, Number of run goroutines 100
Job number: 218; Counter: 200, Number of run goroutines 100
Job number: 222; Counter: 200, Number of run goroutines 100
Job number: 220; Counter: 200, Number of run goroutines 100
Job number: 221; Counter: 200, Number of run goroutines 100
Job number: 225; Counter: 200, Number of run goroutines 100
Job number: 226; Counter: 200, Number of run goroutines 100
Job number: 224; Counter: 200, Number of run goroutines 100
Job number: 229; Counter: 200, Number of run goroutines 100
Job number: 230; Counter: 200, Number of run goroutines 100
Job number: 228; Counter: 200, Number of run goroutines 100
Job number: 235; Counter: 200, Number of run goroutines 100
Job number: 242; Counter: 200, Number of run goroutines 100
Job number: 249; Counter: 200, Number of run goroutines 100
Job number: 298; Counter: 200, Number of run goroutines 100
Job number: 211; Counter: 200, Number of run goroutines 100
Job number: 212; Counter: 200, Number of run goroutines 100
Job number: 300; Counter: 200, Number of run goroutines 100
Job number: 213; Counter: 200, Number of run goroutines 100
Job number: 219; Counter: 200, Number of run goroutines 100
Job number: 207; Counter: 200, Number of run goroutines 100
Job number: 202; Counter: 200, Number of run goroutines 100
Job number: 239; Counter: 200, Number of run goroutines 100
Job number: 251; Counter: 200, Number of run goroutines 100
Job number: 223; Counter: 200, Number of run goroutines 100
Job number: 210; Counter: 200, Number of run goroutines 100
Job number: 275; Counter: 200, Number of run goroutines 100
Job number: 232; Counter: 200, Number of run goroutines 100
Job number: 227; Counter: 200, Number of run goroutines 100
Job number: 231; Counter: 200, Number of run goroutines 100
Job number: 264; Counter: 200, Number of run goroutines 100
Job number: 252; Counter: 200, Number of run goroutines 100
Job number: 276; Counter: 200, Number of run goroutines 100
Job number: 277; Counter: 200, Number of run goroutines 100
Job number: 253; Counter: 200, Number of run goroutines 100
Job number: 254; Counter: 200, Number of run goroutines 100
Job number: 278; Counter: 200, Number of run goroutines 100
Job number: 255; Counter: 200, Number of run goroutines 100
Job number: 280; Counter: 200, Number of run goroutines 100
Job number: 267; Counter: 200, Number of run goroutines 100
Job number: 256; Counter: 200, Number of run goroutines 100
Job number: 236; Counter: 200, Number of run goroutines 100
Job number: 281; Counter: 200, Number of run goroutines 100
Job number: 204; Counter: 200, Number of run goroutines 100
Job number: 257; Counter: 200, Number of run goroutines 100
Job number: 279; Counter: 200, Number of run goroutines 100
Job number: 258; Counter: 200, Number of run goroutines 100
Job number: 205; Counter: 200, Number of run goroutines 100
Job number: 259; Counter: 200, Number of run goroutines 100
Job number: 203; Counter: 200, Number of run goroutines 100
Job number: 260; Counter: 200, Number of run goroutines 100
Job number: 268; Counter: 200, Number of run goroutines 100
Job number: 261; Counter: 200, Number of run goroutines 100
Job number: 265; Counter: 200, Number of run goroutines 100
Job number: 262; Counter: 200, Number of run goroutines 100
Job number: 282; Counter: 200, Number of run goroutines 100
Job number: 263; Counter: 200, Number of run goroutines 100
Job number: 283; Counter: 200, Number of run goroutines 100
Job number: 206; Counter: 200, Number of run goroutines 100
Job number: 271; Counter: 200, Number of run goroutines 100
Job number: 285; Counter: 200, Number of run goroutines 100
Job number: 208; Counter: 200, Number of run goroutines 100
Job number: 266; Counter: 200, Number of run goroutines 100
Job number: 284; Counter: 200, Number of run goroutines 100
Job number: 269; Counter: 200, Number of run goroutines 100
Job number: 287; Counter: 200, Number of run goroutines 100
Job number: 209; Counter: 200, Number of run goroutines 100
Job number: 272; Counter: 200, Number of run goroutines 100
Job number: 270; Counter: 200, Number of run goroutines 100
Job number: 288; Counter: 200, Number of run goroutines 100
Job number: 233; Counter: 200, Number of run goroutines 100
Job number: 286; Counter: 200, Number of run goroutines 100
Job number: 234; Counter: 200, Number of run goroutines 100
Job number: 273; Counter: 200, Number of run goroutines 100
Job number: 292; Counter: 200, Number of run goroutines 100
Job number: 243; Counter: 200, Number of run goroutines 100
Job number: 274; Counter: 200, Number of run goroutines 100
Job number: 237; Counter: 200, Number of run goroutines 100
Job number: 238; Counter: 200, Number of run goroutines 100
Job number: 296; Counter: 200, Number of run goroutines 100
Job number: 293; Counter: 200, Number of run goroutines 100
Job number: 241; Counter: 200, Number of run goroutines 100
Job number: 291; Counter: 200, Number of run goroutines 100
Job number: 245; Counter: 200, Number of run goroutines 100
Job number: 290; Counter: 200, Number of run goroutines 100
Job number: 246; Counter: 200, Number of run goroutines 100
Job number: 244; Counter: 200, Number of run goroutines 100
Job number: 240; Counter: 200, Number of run goroutines 100
Job number: 247; Counter: 200, Number of run goroutines 100
Job number: 248; Counter: 200, Number of run goroutines 100
Job number: 295; Counter: 200, Number of run goroutines 100
Job number: 289; Counter: 200, Number of run goroutines 100
Job number: 294; Counter: 200, Number of run goroutines 100
Job number: 297; Counter: 200, Number of run goroutines 100
Job number: 345; Counter: 300, Number of run goroutines 100 
Job number: 352; Counter: 300, Number of run goroutines 100 
Job number: 301; Counter: 300, Number of run goroutines 100
Job number: 312; Counter: 300, Number of run goroutines 100
Job number: 313; Counter: 300, Number of run goroutines 100
Job number: 314; Counter: 300, Number of run goroutines 100
Job number: 316; Counter: 300, Number of run goroutines 100
Job number: 317; Counter: 300, Number of run goroutines 100
Job number: 315; Counter: 300, Number of run goroutines 100
Job number: 318; Counter: 300, Number of run goroutines 100
Job number: 319; Counter: 300, Number of run goroutines 100
Job number: 320; Counter: 300, Number of run goroutines 100
Job number: 321; Counter: 300, Number of run goroutines 100
Job number: 322; Counter: 300, Number of run goroutines 100
Job number: 337; Counter: 300, Number of run goroutines 100
Job number: 333; Counter: 300, Number of run goroutines 100
Job number: 376; Counter: 316, Number of run goroutines 100
Job number: 334; Counter: 300, Number of run goroutines 100
Job number: 346; Counter: 300, Number of run goroutines 100
Job number: 303; Counter: 300, Number of run goroutines 100
Job number: 347; Counter: 300, Number of run goroutines 100
Job number: 304; Counter: 300, Number of run goroutines 100
Job number: 348; Counter: 300, Number of run goroutines 100
Job number: 305; Counter: 300, Number of run goroutines 100
Job number: 379; Counter: 324, Number of run goroutines 100
Job number: 328; Counter: 300, Number of run goroutines 100
Job number: 353; Counter: 300, Number of run goroutines 100
Job number: 381; Counter: 327, Number of run goroutines 100
Job number: 354; Counter: 300, Number of run goroutines 100
Job number: 327; Counter: 300, Number of run goroutines 100
Job number: 350; Counter: 300, Number of run goroutines 100
Job number: 390; Counter: 331, Number of run goroutines 100
Job number: 351; Counter: 300, Number of run goroutines 100
Job number: 331; Counter: 300, Number of run goroutines 100
Job number: 309; Counter: 300, Number of run goroutines 100
Job number: 336; Counter: 300, Number of run goroutines 100
Job number: 391; Counter: 336, Number of run goroutines 100
Job number: 338; Counter: 300, Number of run goroutines 100
Job number: 306; Counter: 300, Number of run goroutines 100
Job number: 340; Counter: 300, Number of run goroutines 100
Job number: 395; Counter: 340, Number of run goroutines 100
Job number: 311; Counter: 300, Number of run goroutines 100
Job number: 308; Counter: 300, Number of run goroutines 100
Job number: 332; Counter: 300, Number of run goroutines 100
Job number: 339; Counter: 300, Number of run goroutines 100
Job number: 323; Counter: 300, Number of run goroutines 100
Job number: 324; Counter: 300, Number of run goroutines 100
Job number: 310; Counter: 300, Number of run goroutines 100
Job number: 342; Counter: 300, Number of run goroutines 100
Job number: 343; Counter: 300, Number of run goroutines 100
Job number: 341; Counter: 300, Number of run goroutines 100
Job number: 325; Counter: 300, Number of run goroutines 100
Job number: 344; Counter: 300, Number of run goroutines 100
Job number: 326; Counter: 300, Number of run goroutines 100
Job number: 358; Counter: 305, Number of run goroutines 100
Job number: 355; Counter: 305, Number of run goroutines 100
Job number: 356; Counter: 305, Number of run goroutines 100
Job number: 357; Counter: 305, Number of run goroutines 100
Job number: 362; Counter: 306, Number of run goroutines 100
Job number: 359; Counter: 306, Number of run goroutines 100
Job number: 360; Counter: 306, Number of run goroutines 100
Job number: 361; Counter: 306, Number of run goroutines 100
Job number: 363; Counter: 306, Number of run goroutines 100
Job number: 365; Counter: 311, Number of run goroutines 100
Job number: 368; Counter: 311, Number of run goroutines 100
Job number: 366; Counter: 311, Number of run goroutines 100
Job number: 367; Counter: 311, Number of run goroutines 100
Job number: 364; Counter: 311, Number of run goroutines 100
Job number: 369; Counter: 316, Number of run goroutines 100
Job number: 302; Counter: 300, Number of run goroutines 100
Job number: 377; Counter: 316, Number of run goroutines 100
Job number: 374; Counter: 316, Number of run goroutines 100
Job number: 372; Counter: 316, Number of run goroutines 100
Job number: 375; Counter: 316, Number of run goroutines 100
Job number: 370; Counter: 316, Number of run goroutines 100
Job number: 373; Counter: 316, Number of run goroutines 100
Job number: 371; Counter: 316, Number of run goroutines 100
Job number: 349; Counter: 300, Number of run goroutines 100
Job number: 378; Counter: 324, Number of run goroutines 100
Job number: 380; Counter: 324, Number of run goroutines 100
Job number: 383; Counter: 327, Number of run goroutines 100
Job number: 329; Counter: 300, Number of run goroutines 100
Job number: 384; Counter: 327, Number of run goroutines 100
Job number: 382; Counter: 327, Number of run goroutines 100
Job number: 385; Counter: 327, Number of run goroutines 100
Job number: 389; Counter: 331, Number of run goroutines 100
Job number: 330; Counter: 300, Number of run goroutines 100
Job number: 386; Counter: 331, Number of run goroutines 100
Job number: 388; Counter: 331, Number of run goroutines 100
Job number: 387; Counter: 331, Number of run goroutines 100
Job number: 393; Counter: 336, Number of run goroutines 100
Job number: 335; Counter: 300, Number of run goroutines 100
Job number: 392; Counter: 336, Number of run goroutines 100
Job number: 394; Counter: 336, Number of run goroutines 100
Job number: 307; Counter: 300, Number of run goroutines 100
Job number: 397; Counter: 340, Number of run goroutines 100
Job number: 396; Counter: 340, Number of run goroutines 100
Job number: 398; Counter: 345, Number of run goroutines 100
Job number: 400; Counter: 345, Number of run goroutines 100
Job number: 399; Counter: 345, Number of run goroutines 100
Job number: 454; Counter: 400, Number of run goroutines 100 
Job number: 460; Counter: 400, Number of run goroutines 100 
Job number: 402; Counter: 400, Number of run goroutines 100
Job number: 403; Counter: 400, Number of run goroutines 100
Job number: 404; Counter: 400, Number of run goroutines 100
Job number: 405; Counter: 400, Number of run goroutines 100
Job number: 407; Counter: 400, Number of run goroutines 100
Job number: 406; Counter: 400, Number of run goroutines 100
Job number: 408; Counter: 400, Number of run goroutines 100
Job number: 409; Counter: 400, Number of run goroutines 100
Job number: 410; Counter: 400, Number of run goroutines 100
Job number: 416; Counter: 400, Number of run goroutines 100
Job number: 424; Counter: 400, Number of run goroutines 100
Job number: 436; Counter: 400, Number of run goroutines 100
Job number: 448; Counter: 400, Number of run goroutines 100
Job number: 427; Counter: 400, Number of run goroutines 100
Job number: 445; Counter: 400, Number of run goroutines 100
Job number: 446; Counter: 400, Number of run goroutines 100
Job number: 458; Counter: 400, Number of run goroutines 100
Job number: 447; Counter: 400, Number of run goroutines 100
Job number: 411; Counter: 400, Number of run goroutines 100
Job number: 412; Counter: 400, Number of run goroutines 100
Job number: 414; Counter: 400, Number of run goroutines 100
Job number: 415; Counter: 400, Number of run goroutines 100
Job number: 413; Counter: 400, Number of run goroutines 100
Job number: 451; Counter: 400, Number of run goroutines 100
Job number: 418; Counter: 400, Number of run goroutines 100
Job number: 419; Counter: 400, Number of run goroutines 100
Job number: 449; Counter: 400, Number of run goroutines 100
Job number: 440; Counter: 400, Number of run goroutines 100
Job number: 417; Counter: 400, Number of run goroutines 100
Job number: 450; Counter: 400, Number of run goroutines 100
Job number: 421; Counter: 400, Number of run goroutines 100
Job number: 431; Counter: 400, Number of run goroutines 100
Job number: 433; Counter: 400, Number of run goroutines 100
Job number: 437; Counter: 400, Number of run goroutines 100
Job number: 423; Counter: 400, Number of run goroutines 100
Job number: 443; Counter: 400, Number of run goroutines 100
Job number: 420; Counter: 400, Number of run goroutines 100
Job number: 441; Counter: 400, Number of run goroutines 100
Job number: 456; Counter: 400, Number of run goroutines 100
Job number: 455; Counter: 400, Number of run goroutines 100
Job number: 439; Counter: 400, Number of run goroutines 100
Job number: 457; Counter: 400, Number of run goroutines 100
Job number: 422; Counter: 400, Number of run goroutines 100
Job number: 459; Counter: 400, Number of run goroutines 100
Job number: 426; Counter: 400, Number of run goroutines 100
Job number: 438; Counter: 400, Number of run goroutines 100
Job number: 429; Counter: 400, Number of run goroutines 100
Job number: 432; Counter: 400, Number of run goroutines 100
Job number: 442; Counter: 400, Number of run goroutines 100
Job number: 401; Counter: 400, Number of run goroutines 100
Job number: 435; Counter: 400, Number of run goroutines 100
Job number: 425; Counter: 400, Number of run goroutines 100
Job number: 434; Counter: 400, Number of run goroutines 100
Job number: 444; Counter: 400, Number of run goroutines 100
Job number: 428; Counter: 400, Number of run goroutines 100
Job number: 453; Counter: 400, Number of run goroutines 100
Job number: 452; Counter: 400, Number of run goroutines 100
Job number: 461; Counter: 400, Number of run goroutines 100
Job number: 430; Counter: 400, Number of run goroutines 100
Job number: 463; Counter: 406, Number of run goroutines 100
Job number: 465; Counter: 406, Number of run goroutines 100
Job number: 464; Counter: 406, Number of run goroutines 100
Job number: 466; Counter: 406, Number of run goroutines 100
Job number: 462; Counter: 406, Number of run goroutines 100
Job number: 467; Counter: 410, Number of run goroutines 100
Job number: 469; Counter: 410, Number of run goroutines 100
Job number: 470; Counter: 410, Number of run goroutines 100
Job number: 468; Counter: 410, Number of run goroutines 100
Job number: 474; Counter: 415, Number of run goroutines 100
Job number: 473; Counter: 415, Number of run goroutines 100
Job number: 471; Counter: 415, Number of run goroutines 100
Job number: 472; Counter: 415, Number of run goroutines 100
Job number: 475; Counter: 415, Number of run goroutines 100
Job number: 476; Counter: 419, Number of run goroutines 100
Job number: 478; Counter: 419, Number of run goroutines 100
Job number: 479; Counter: 419, Number of run goroutines 100
Job number: 477; Counter: 419, Number of run goroutines 100
Job number: 483; Counter: 423, Number of run goroutines 100
Job number: 480; Counter: 423, Number of run goroutines 100
Job number: 484; Counter: 423, Number of run goroutines 100
Job number: 481; Counter: 423, Number of run goroutines 100
Job number: 482; Counter: 423, Number of run goroutines 100
Job number: 485; Counter: 426, Number of run goroutines 100
Job number: 486; Counter: 426, Number of run goroutines 100
Job number: 488; Counter: 426, Number of run goroutines 100
Job number: 487; Counter: 426, Number of run goroutines 100
Job number: 490; Counter: 431, Number of run goroutines 100
Job number: 493; Counter: 431, Number of run goroutines 100
Job number: 491; Counter: 431, Number of run goroutines 100
Job number: 492; Counter: 431, Number of run goroutines 100
Job number: 489; Counter: 431, Number of run goroutines 100
Job number: 498; Counter: 436, Number of run goroutines 100
Job number: 494; Counter: 436, Number of run goroutines 100
Job number: 496; Counter: 436, Number of run goroutines 100
Job number: 495; Counter: 436, Number of run goroutines 100
Job number: 497; Counter: 436, Number of run goroutines 100
Job number: 500; Counter: 437, Number of run goroutines 100
Job number: 499; Counter: 437, Number of run goroutines 100
Job number: 563; Counter: 500, Number of run goroutines 100 
Job number: 518; Counter: 500, Number of run goroutines 100 
Job number: 503; Counter: 500, Number of run goroutines 100
Job number: 504; Counter: 500, Number of run goroutines 100
Job number: 505; Counter: 500, Number of run goroutines 100
Job number: 508; Counter: 500, Number of run goroutines 100
Job number: 509; Counter: 500, Number of run goroutines 100
Job number: 506; Counter: 500, Number of run goroutines 100
Job number: 507; Counter: 500, Number of run goroutines 100
Job number: 510; Counter: 500, Number of run goroutines 100
Job number: 511; Counter: 500, Number of run goroutines 100
Job number: 512; Counter: 500, Number of run goroutines 100
Job number: 513; Counter: 500, Number of run goroutines 100
Job number: 514; Counter: 500, Number of run goroutines 100
Job number: 515; Counter: 500, Number of run goroutines 100
Job number: 536; Counter: 500, Number of run goroutines 100
Job number: 519; Counter: 500, Number of run goroutines 100
Job number: 520; Counter: 500, Number of run goroutines 100
Job number: 516; Counter: 500, Number of run goroutines 100
Job number: 521; Counter: 500, Number of run goroutines 100
Job number: 517; Counter: 500, Number of run goroutines 100
Job number: 522; Counter: 500, Number of run goroutines 100
Job number: 524; Counter: 500, Number of run goroutines 100
Job number: 525; Counter: 500, Number of run goroutines 100
Job number: 548; Counter: 500, Number of run goroutines 100
Job number: 557; Counter: 500, Number of run goroutines 100
Job number: 523; Counter: 500, Number of run goroutines 100
Job number: 502; Counter: 500, Number of run goroutines 100
Job number: 549; Counter: 500, Number of run goroutines 100
Job number: 526; Counter: 500, Number of run goroutines 100
Job number: 501; Counter: 500, Number of run goroutines 100
Job number: 550; Counter: 500, Number of run goroutines 100
Job number: 527; Counter: 500, Number of run goroutines 100
Job number: 540; Counter: 500, Number of run goroutines 100
Job number: 528; Counter: 500, Number of run goroutines 100
Job number: 537; Counter: 500, Number of run goroutines 100
Job number: 529; Counter: 500, Number of run goroutines 100
Job number: 551; Counter: 500, Number of run goroutines 100
Job number: 538; Counter: 500, Number of run goroutines 100
Job number: 530; Counter: 500, Number of run goroutines 100
Job number: 552; Counter: 500, Number of run goroutines 100
Job number: 531; Counter: 500, Number of run goroutines 100
Job number: 539; Counter: 500, Number of run goroutines 100
Job number: 553; Counter: 500, Number of run goroutines 100
Job number: 532; Counter: 500, Number of run goroutines 100
Job number: 533; Counter: 500, Number of run goroutines 100
Job number: 541; Counter: 500, Number of run goroutines 100
Job number: 534; Counter: 500, Number of run goroutines 100
Job number: 556; Counter: 500, Number of run goroutines 100
Job number: 535; Counter: 500, Number of run goroutines 100
Job number: 560; Counter: 500, Number of run goroutines 100
Job number: 545; Counter: 500, Number of run goroutines 100
Job number: 555; Counter: 500, Number of run goroutines 100
Job number: 543; Counter: 500, Number of run goroutines 100
Job number: 558; Counter: 500, Number of run goroutines 100
Job number: 544; Counter: 500, Number of run goroutines 100
Job number: 554; Counter: 500, Number of run goroutines 100
Job number: 546; Counter: 500, Number of run goroutines 100
Job number: 559; Counter: 500, Number of run goroutines 100
Job number: 562; Counter: 500, Number of run goroutines 100
Job number: 561; Counter: 500, Number of run goroutines 100
Job number: 547; Counter: 500, Number of run goroutines 100
Job number: 565; Counter: 500, Number of run goroutines 100
Job number: 564; Counter: 500, Number of run goroutines 100
Job number: 566; Counter: 500, Number of run goroutines 100
Job number: 567; Counter: 500, Number of run goroutines 100
Job number: 568; Counter: 500, Number of run goroutines 100
Job number: 542; Counter: 500, Number of run goroutines 100
Job number: 569; Counter: 506, Number of run goroutines 100
Job number: 572; Counter: 506, Number of run goroutines 100
Job number: 573; Counter: 506, Number of run goroutines 100
Job number: 571; Counter: 506, Number of run goroutines 100
Job number: 570; Counter: 506, Number of run goroutines 100
Job number: 574; Counter: 509, Number of run goroutines 100
Job number: 577; Counter: 509, Number of run goroutines 100
Job number: 575; Counter: 509, Number of run goroutines 100
Job number: 576; Counter: 509, Number of run goroutines 100
Job number: 581; Counter: 512, Number of run goroutines 100
Job number: 578; Counter: 512, Number of run goroutines 100
Job number: 579; Counter: 512, Number of run goroutines 100
Job number: 580; Counter: 512, Number of run goroutines 100
Job number: 582; Counter: 516, Number of run goroutines 100
Job number: 584; Counter: 516, Number of run goroutines 100
Job number: 585; Counter: 516, Number of run goroutines 100
Job number: 583; Counter: 516, Number of run goroutines 100
Job number: 586; Counter: 516, Number of run goroutines 100
Job number: 589; Counter: 520, Number of run goroutines 100
Job number: 590; Counter: 520, Number of run goroutines 100
Job number: 591; Counter: 520, Number of run goroutines 100
Job number: 588; Counter: 520, Number of run goroutines 100
Job number: 587; Counter: 520, Number of run goroutines 100
Job number: 595; Counter: 523, Number of run goroutines 100
Job number: 593; Counter: 523, Number of run goroutines 100
Job number: 594; Counter: 523, Number of run goroutines 100
Job number: 592; Counter: 523, Number of run goroutines 100
Job number: 596; Counter: 523, Number of run goroutines 100
Job number: 598; Counter: 527, Number of run goroutines 100
Job number: 600; Counter: 527, Number of run goroutines 100
Job number: 599; Counter: 527, Number of run goroutines 100
Job number: 597; Counter: 527, Number of run goroutines 100
Job number: 662; Counter: 600, Number of run goroutines 100 
Job number: 654; Counter: 600, Number of run goroutines 100 
Job number: 667; Counter: 602, Number of run goroutines 100
Job number: 671; Counter: 603, Number of run goroutines 99
Job number: 656; Counter: 600, Number of run goroutines 100
Job number: 677; Counter: 605, Number of run goroutines 97
Job number: 653; Counter: 600, Number of run goroutines 100
Job number: 659; Counter: 600, Number of run goroutines 100
Job number: 655; Counter: 600, Number of run goroutines 100
Job number: 683; Counter: 609, Number of run goroutines 100
Job number: 691; Counter: 610, Number of run goroutines 100
Job number: 602; Counter: 600, Number of run goroutines 100
Job number: 695; Counter: 612, Number of run goroutines 100
Job number: 604; Counter: 600, Number of run goroutines 100
Job number: 700; Counter: 614, Number of run goroutines 100
Job number: 608; Counter: 600, Number of run goroutines 100
Job number: 609; Counter: 600, Number of run goroutines 100
Job number: 610; Counter: 600, Number of run goroutines 100
Job number: 611; Counter: 600, Number of run goroutines 100
Job number: 612; Counter: 600, Number of run goroutines 100
Job number: 613; Counter: 600, Number of run goroutines 100
Job number: 614; Counter: 600, Number of run goroutines 100
Job number: 619; Counter: 600, Number of run goroutines 100
Job number: 620; Counter: 600, Number of run goroutines 100
Job number: 621; Counter: 600, Number of run goroutines 100
Job number: 622; Counter: 600, Number of run goroutines 100
Job number: 605; Counter: 600, Number of run goroutines 100
Job number: 623; Counter: 600, Number of run goroutines 100
Job number: 624; Counter: 600, Number of run goroutines 100
Job number: 606; Counter: 600, Number of run goroutines 100
Job number: 618; Counter: 600, Number of run goroutines 100
Job number: 635; Counter: 600, Number of run goroutines 100
Job number: 601; Counter: 600, Number of run goroutines 100
Job number: 625; Counter: 600, Number of run goroutines 100
Job number: 615; Counter: 600, Number of run goroutines 100
Job number: 617; Counter: 600, Number of run goroutines 100
Job number: 639; Counter: 600, Number of run goroutines 100
Job number: 626; Counter: 600, Number of run goroutines 100
Job number: 616; Counter: 600, Number of run goroutines 100
Job number: 644; Counter: 600, Number of run goroutines 100
Job number: 630; Counter: 600, Number of run goroutines 100
Job number: 645; Counter: 600, Number of run goroutines 100
Job number: 633; Counter: 600, Number of run goroutines 100
Job number: 629; Counter: 600, Number of run goroutines 100
Job number: 647; Counter: 600, Number of run goroutines 100
Job number: 631; Counter: 600, Number of run goroutines 100
Job number: 627; Counter: 600, Number of run goroutines 100
Job number: 646; Counter: 600, Number of run goroutines 100
Job number: 634; Counter: 600, Number of run goroutines 100
Job number: 649; Counter: 600, Number of run goroutines 100
Job number: 650; Counter: 600, Number of run goroutines 100
Job number: 638; Counter: 600, Number of run goroutines 100
Job number: 632; Counter: 600, Number of run goroutines 100
Job number: 636; Counter: 600, Number of run goroutines 100
Job number: 660; Counter: 600, Number of run goroutines 100
Job number: 661; Counter: 600, Number of run goroutines 100
Job number: 628; Counter: 600, Number of run goroutines 100
Job number: 642; Counter: 600, Number of run goroutines 100
Job number: 637; Counter: 600, Number of run goroutines 100
Job number: 640; Counter: 600, Number of run goroutines 100
Job number: 641; Counter: 600, Number of run goroutines 100
Job number: 643; Counter: 606, Number of run goroutines 96
Job number: 672; Counter: 606, Number of run goroutines 96
Job number: 651; Counter: 600, Number of run goroutines 100
Job number: 663; Counter: 606, Number of run goroutines 100
Job number: 664; Counter: 606, Number of run goroutines 100
Job number: 665; Counter: 606, Number of run goroutines 100
Job number: 666; Counter: 606, Number of run goroutines 100
Job number: 652; Counter: 600, Number of run goroutines 100
Job number: 657; Counter: 600, Number of run goroutines 100
Job number: 668; Counter: 606, Number of run goroutines 100
Job number: 669; Counter: 606, Number of run goroutines 100
Job number: 670; Counter: 606, Number of run goroutines 100
Job number: 674; Counter: 606, Number of run goroutines 100
Job number: 673; Counter: 606, Number of run goroutines 100
Job number: 675; Counter: 606, Number of run goroutines 100
Job number: 676; Counter: 606, Number of run goroutines 100
Job number: 682; Counter: 606, Number of run goroutines 100
Job number: 679; Counter: 606, Number of run goroutines 100
Job number: 678; Counter: 606, Number of run goroutines 100
Job number: 680; Counter: 606, Number of run goroutines 100 
Job number: 681; Counter: 606, Number of run goroutines 100
Job number: 686; Counter: 609, Number of run goroutines 100
Job number: 684; Counter: 609, Number of run goroutines 100
Job number: 687; Counter: 609, Number of run goroutines 100
Job number: 658; Counter: 600, Number of run goroutines 100
Job number: 685; Counter: 609, Number of run goroutines 100
Job number: 689; Counter: 610, Number of run goroutines 100
Job number: 607; Counter: 600, Number of run goroutines 100
Job number: 693; Counter: 610, Number of run goroutines 100
Job number: 692; Counter: 610, Number of run goroutines 100
Job number: 688; Counter: 610, Number of run goroutines 100
Job number: 690; Counter: 610, Number of run goroutines 100
Job number: 698; Counter: 612, Number of run goroutines 100
Job number: 603; Counter: 600, Number of run goroutines 100
Job number: 697; Counter: 612, Number of run goroutines 100
Job number: 696; Counter: 612, Number of run goroutines 100
Job number: 694; Counter: 612, Number of run goroutines 100
Job number: 648; Counter: 600, Number of run goroutines 100
Job number: 699; Counter: 614, Number of run goroutines 100
Job number: 712; Counter: 700, Number of run goroutines 100 
Job number: 720; Counter: 700, Number of run goroutines 100 
Job number: 721; Counter: 700, Number of run goroutines 100
Job number: 701; Counter: 700, Number of run goroutines 100
Job number: 728; Counter: 704, Number of run goroutines 100
Job number: 730; Counter: 705, Number of run goroutines 100
Job number: 731; Counter: 706, Number of run goroutines 100
Job number: 738; Counter: 707, Number of run goroutines 100
Job number: 743; Counter: 708, Number of run goroutines 100
Job number: 746; Counter: 709, Number of run goroutines 100
Job number: 750; Counter: 710, Number of run goroutines 100
Job number: 751; Counter: 711, Number of run goroutines 100
Job number: 758; Counter: 712, Number of run goroutines 100
Job number: 759; Counter: 713, Number of run goroutines 100
Job number: 702; Counter: 700, Number of run goroutines 100
Job number: 719; Counter: 700, Number of run goroutines 100
Job number: 703; Counter: 700, Number of run goroutines 100
Job number: 704; Counter: 700, Number of run goroutines 100
Job number: 709; Counter: 700, Number of run goroutines 100
Job number: 784; Counter: 719, Number of run goroutines 100
Job number: 710; Counter: 700, Number of run goroutines 100
Job number: 723; Counter: 701, Number of run goroutines 100
Job number: 722; Counter: 701, Number of run goroutines 100
Job number: 725; Counter: 702, Number of run goroutines 100
Job number: 724; Counter: 702, Number of run goroutines 100
Job number: 727; Counter: 703, Number of run goroutines 100
Job number: 726; Counter: 703, Number of run goroutines 100
Job number: 707; Counter: 700, Number of run goroutines 100
Job number: 708; Counter: 700, Number of run goroutines 100
Job number: 729; Counter: 705, Number of run goroutines 100
Job number: 711; Counter: 700, Number of run goroutines 100 
Job number: 732; Counter: 706, Number of run goroutines 100
Job number: 733; Counter: 706, Number of run goroutines 100
Job number: 734; Counter: 707, Number of run goroutines 100
Job number: 713; Counter: 700, Number of run goroutines 100
Job number: 736; Counter: 707, Number of run goroutines 100
Job number: 735; Counter: 707, Number of run goroutines 100
Job number: 737; Counter: 707, Number of run goroutines 100
Job number: 739; Counter: 708, Number of run goroutines 100
Job number: 715; Counter: 700, Number of run goroutines 100
Job number: 741; Counter: 708, Number of run goroutines 100
Job number: 740; Counter: 708, Number of run goroutines 100
Job number: 742; Counter: 708, Number of run goroutines 100
Job number: 744; Counter: 709, Number of run goroutines 100
Job number: 716; Counter: 700, Number of run goroutines 100
Job number: 745; Counter: 709, Number of run goroutines 100
Job number: 747; Counter: 709, Number of run goroutines 100
Job number: 714; Counter: 700, Number of run goroutines 100
Job number: 748; Counter: 710, Number of run goroutines 100
Job number: 749; Counter: 710, Number of run goroutines 100
Job number: 752; Counter: 711, Number of run goroutines 100
Job number: 706; Counter: 700, Number of run goroutines 100
Job number: 753; Counter: 711, Number of run goroutines 100
Job number: 754; Counter: 711, Number of run goroutines 100
Job number: 756; Counter: 712, Number of run goroutines 100
Job number: 717; Counter: 700, Number of run goroutines 100
Job number: 755; Counter: 712, Number of run goroutines 100
Job number: 757; Counter: 712, Number of run goroutines 100
Job number: 718; Counter: 700, Number of run goroutines 100
Job number: 761; Counter: 713, Number of run goroutines 100
Job number: 760; Counter: 713, Number of run goroutines 100
Job number: 762; Counter: 713, Number of run goroutines 100
Job number: 766; Counter: 714, Number of run goroutines 100
Job number: 763; Counter: 714, Number of run goroutines 100
Job number: 764; Counter: 714, Number of run goroutines 100
Job number: 765; Counter: 714, Number of run goroutines 100
Job number: 769; Counter: 715, Number of run goroutines 100
Job number: 767; Counter: 715, Number of run goroutines 100
Job number: 768; Counter: 715, Number of run goroutines 100
Job number: 770; Counter: 715, Number of run goroutines 100 
Job number: 772; Counter: 716, Number of run goroutines 100
Job number: 774; Counter: 716, Number of run goroutines 100
Job number: 773; Counter: 716, Number of run goroutines 100
Job number: 771; Counter: 716, Number of run goroutines 100
Job number: 775; Counter: 717, Number of run goroutines 100
Job number: 776; Counter: 717, Number of run goroutines 100
Job number: 777; Counter: 717, Number of run goroutines 100
Job number: 778; Counter: 717, Number of run goroutines 100
Job number: 779; Counter: 717, Number of run goroutines 100
Job number: 781; Counter: 718, Number of run goroutines 100
Job number: 780; Counter: 718, Number of run goroutines 100
Job number: 705; Counter: 700, Number of run goroutines 100
Job number: 783; Counter: 719, Number of run goroutines 100
Job number: 785; Counter: 719, Number of run goroutines 100
Job number: 782; Counter: 719, Number of run goroutines 100
Job number: 786; Counter: 720, Number of run goroutines 100
Job number: 787; Counter: 720, Number of run goroutines 100
Job number: 790; Counter: 720, Number of run goroutines 100
Job number: 788; Counter: 720, Number of run goroutines 100
Job number: 789; Counter: 720, Number of run goroutines 100
Job number: 797; Counter: 722, Number of run goroutines 100
Job number: 792; Counter: 722, Number of run goroutines 100
Job number: 793; Counter: 722, Number of run goroutines 100
Job number: 794; Counter: 722, Number of run goroutines 100
Job number: 791; Counter: 722, Number of run goroutines 100
Job number: 798; Counter: 722, Number of run goroutines 100
Job number: 795; Counter: 722, Number of run goroutines 100
Job number: 796; Counter: 722, Number of run goroutines 100
Job number: 799; Counter: 723, Number of run goroutines 100
Job number: 800; Counter: 723, Number of run goroutines 100
Job number: 812; Counter: 800, Number of run goroutines 100 
Job number: 811; Counter: 800, Number of run goroutines 100 
Job number: 819; Counter: 801, Number of run goroutines 100
Job number: 810; Counter: 800, Number of run goroutines 100
Job number: 807; Counter: 800, Number of run goroutines 100
Job number: 805; Counter: 800, Number of run goroutines 100
Job number: 808; Counter: 800, Number of run goroutines 100
Job number: 803; Counter: 800, Number of run goroutines 100
Job number: 801; Counter: 800, Number of run goroutines 100
Job number: 802; Counter: 800, Number of run goroutines 100
Job number: 809; Counter: 800, Number of run goroutines 100
Job number: 815; Counter: 800, Number of run goroutines 100
Job number: 814; Counter: 800, Number of run goroutines 100
Job number: 817; Counter: 800, Number of run goroutines 100
Job number: 816; Counter: 800, Number of run goroutines 100
Job number: 804; Counter: 800, Number of run goroutines 100
Job number: 813; Counter: 800, Number of run goroutines 100
Job number: 818; Counter: 801, Number of run goroutines 100
Job number: 806; Counter: 800, Number of run goroutines 100
Job number: 820; Counter: 803, Number of run goroutines 100
Job number: 821; Counter: 804, Number of run goroutines 100
Job number: 822; Counter: 805, Number of run goroutines 100
Job number: 823; Counter: 806, Number of run goroutines 100
Job number: 824; Counter: 807, Number of run goroutines 100
Job number: 825; Counter: 808, Number of run goroutines 100
Job number: 826; Counter: 809, Number of run goroutines 100
Job number: 827; Counter: 810, Number of run goroutines 100
Job number: 828; Counter: 811, Number of run goroutines 100
Job number: 829; Counter: 812, Number of run goroutines 100
Job number: 830; Counter: 813, Number of run goroutines 100 
Job number: 831; Counter: 814, Number of run goroutines 100
Job number: 832; Counter: 815, Number of run goroutines 100
Job number: 852; Counter: 832, Number of run goroutines 100
Job number: 854; Counter: 833, Number of run goroutines 100
Job number: 855; Counter: 834, Number of run goroutines 100
Job number: 856; Counter: 835, Number of run goroutines 100
Job number: 858; Counter: 836, Number of run goroutines 100
Job number: 860; Counter: 837, Number of run goroutines 100
Job number: 862; Counter: 838, Number of run goroutines 100
Job number: 863; Counter: 839, Number of run goroutines 100
Job number: 865; Counter: 840, Number of run goroutines 100
Job number: 842; Counter: 825, Number of run goroutines 100
Job number: 843; Counter: 826, Number of run goroutines 100
Job number: 869; Counter: 843, Number of run goroutines 100
Job number: 845; Counter: 828, Number of run goroutines 100
Job number: 846; Counter: 828, Number of run goroutines 100
Job number: 847; Counter: 829, Number of run goroutines 100
Job number: 848; Counter: 829, Number of run goroutines 100
Job number: 850; Counter: 830, Number of run goroutines 100
Job number: 849; Counter: 830, Number of run goroutines 100
Job number: 851; Counter: 831, Number of run goroutines 100
Job number: 833; Counter: 816, Number of run goroutines 100
Job number: 834; Counter: 817, Number of run goroutines 100
Job number: 853; Counter: 833, Number of run goroutines 100
Job number: 835; Counter: 818, Number of run goroutines 100
Job number: 836; Counter: 819, Number of run goroutines 100
Job number: 857; Counter: 835, Number of run goroutines 100
Job number: 837; Counter: 820, Number of run goroutines 100
Job number: 859; Counter: 836, Number of run goroutines 100
Job number: 838; Counter: 821, Number of run goroutines 100
Job number: 861; Counter: 837, Number of run goroutines 100
Job number: 839; Counter: 822, Number of run goroutines 100
Job number: 841; Counter: 824, Number of run goroutines 100
Job number: 864; Counter: 839, Number of run goroutines 100
Job number: 840; Counter: 824, Number of run goroutines 100
Job number: 867; Counter: 841, Number of run goroutines 100
Job number: 866; Counter: 841, Number of run goroutines 100
Job number: 871; Counter: 843, Number of run goroutines 100 
Job number: 844; Counter: 827, Number of run goroutines 100
Job number: 870; Counter: 843, Number of run goroutines 100
Job number: 868; Counter: 843, Number of run goroutines 100
Job number: 872; Counter: 845, Number of run goroutines 100
Job number: 873; Counter: 845, Number of run goroutines 100
Job number: 874; Counter: 845, Number of run goroutines 100
Job number: 875; Counter: 845, Number of run goroutines 100
Job number: 878; Counter: 846, Number of run goroutines 100
Job number: 876; Counter: 846, Number of run goroutines 100
Job number: 877; Counter: 846, Number of run goroutines 100
Job number: 881; Counter: 847, Number of run goroutines 100
Job number: 879; Counter: 847, Number of run goroutines 100
Job number: 880; Counter: 847, Number of run goroutines 100
Job number: 885; Counter: 849, Number of run goroutines 100
Job number: 882; Counter: 849, Number of run goroutines 100
Job number: 883; Counter: 849, Number of run goroutines 100
Job number: 884; Counter: 849, Number of run goroutines 100
Job number: 888; Counter: 850, Number of run goroutines 100
Job number: 889; Counter: 850, Number of run goroutines 100
Job number: 886; Counter: 850, Number of run goroutines 100
Job number: 887; Counter: 850, Number of run goroutines 100
Job number: 893; Counter: 852, Number of run goroutines 100
Job number: 890; Counter: 852, Number of run goroutines 100
Job number: 891; Counter: 852, Number of run goroutines 100
Job number: 892; Counter: 852, Number of run goroutines 100
Job number: 897; Counter: 853, Number of run goroutines 100
Job number: 894; Counter: 853, Number of run goroutines 100
Job number: 895; Counter: 853, Number of run goroutines 100
Job number: 896; Counter: 853, Number of run goroutines 100
Job number: 898; Counter: 854, Number of run goroutines 100
Job number: 899; Counter: 854, Number of run goroutines 100
Job number: 900; Counter: 854, Number of run goroutines 100
Job number: 916; Counter: 900, Number of run goroutines 100 
Job number: 909; Counter: 900, Number of run goroutines 100 
Job number: 903; Counter: 900, Number of run goroutines 100
Job number: 904; Counter: 900, Number of run goroutines 100
Job number: 905; Counter: 900, Number of run goroutines 100
Job number: 908; Counter: 900, Number of run goroutines 100
Job number: 906; Counter: 900, Number of run goroutines 100
Job number: 912; Counter: 900, Number of run goroutines 100
Job number: 910; Counter: 900, Number of run goroutines 100
Job number: 911; Counter: 900, Number of run goroutines 100
Job number: 907; Counter: 900, Number of run goroutines 100
Job number: 914; Counter: 900, Number of run goroutines 100
Job number: 915; Counter: 900, Number of run goroutines 100
Job number: 913; Counter: 900, Number of run goroutines 100
Job number: 902; Counter: 900, Number of run goroutines 100
Job number: 901; Counter: 900, Number of run goroutines 100
Job number: 917; Counter: 901, Number of run goroutines 99
Job number: 918; Counter: 903, Number of run goroutines 97
Job number: 919; Counter: 904, Number of run goroutines 96
Job number: 922; Counter: 905, Number of run goroutines 95
Job number: 920; Counter: 905, Number of run goroutines 95
Job number: 921; Counter: 905, Number of run goroutines 95
Job number: 923; Counter: 908, Number of run goroutines 92
Job number: 924; Counter: 909, Number of run goroutines 91
Job number: 934; Counter: 924, Number of run goroutines 76
Job number: 925; Counter: 912, Number of run goroutines 88
Job number: 927; Counter: 913, Number of run goroutines 87
Job number: 928; Counter: 916, Number of run goroutines 84
Job number: 929; Counter: 917, Number of run goroutines 83
Job number: 931; Counter: 920, Number of run goroutines 80
Job number: 930; Counter: 920, Number of run goroutines 80
Job number: 939; Counter: 931, Number of run goroutines 69
Job number: 932; Counter: 922, Number of run goroutines 78
Job number: 926; Counter: 912, Number of run goroutines 88
Job number: 935; Counter: 925, Number of run goroutines 75
Job number: 937; Counter: 928, Number of run goroutines 72
Job number: 936; Counter: 928, Number of run goroutines 72
Job number: 938; Counter: 931, Number of run goroutines 69
Job number: 933; Counter: 922, Number of run goroutines 78
Job number: 940; Counter: 933, Number of run goroutines 67
Job number: 941; Counter: 933, Number of run goroutines 67
Job number: 942; Counter: 935, Number of run goroutines 65 
Job number: 943; Counter: 938, Number of run goroutines 62
Job number: 945; Counter: 938, Number of run goroutines 62
Job number: 944; Counter: 938, Number of run goroutines 62
Job number: 946; Counter: 940, Number of run goroutines 60
Job number: 947; Counter: 940, Number of run goroutines 60
Job number: 948; Counter: 940, Number of run goroutines 60
Job number: 949; Counter: 942, Number of run goroutines 58
Job number: 951; Counter: 944, Number of run goroutines 56
Job number: 950; Counter: 944, Number of run goroutines 56
Job number: 952; Counter: 945, Number of run goroutines 55
Job number: 954; Counter: 946, Number of run goroutines 54
Job number: 953; Counter: 946, Number of run goroutines 54
Job number: 955; Counter: 948, Number of run goroutines 52
Job number: 956; Counter: 949, Number of run goroutines 51
Job number: 959; Counter: 952, Number of run goroutines 48
Job number: 957; Counter: 952, Number of run goroutines 48
Job number: 958; Counter: 952, Number of run goroutines 48
Job number: 960; Counter: 955, Number of run goroutines 45
Job number: 961; Counter: 955, Number of run goroutines 45
Job number: 962; Counter: 955, Number of run goroutines 45
Job number: 964; Counter: 956, Number of run goroutines 44
Job number: 963; Counter: 956, Number of run goroutines 44
Job number: 966; Counter: 958, Number of run goroutines 42
Job number: 965; Counter: 958, Number of run goroutines 42
Job number: 968; Counter: 960, Number of run goroutines 40
Job number: 967; Counter: 960, Number of run goroutines 40
Job number: 969; Counter: 962, Number of run goroutines 38
Job number: 972; Counter: 966, Number of run goroutines 34
Job number: 970; Counter: 966, Number of run goroutines 34
Job number: 971; Counter: 966, Number of run goroutines 34
Job number: 973; Counter: 966, Number of run goroutines 34
Job number: 975; Counter: 968, Number of run goroutines 32
Job number: 974; Counter: 968, Number of run goroutines 32
Job number: 976; Counter: 970, Number of run goroutines 30
Job number: 977; Counter: 970, Number of run goroutines 30
Job number: 980; Counter: 972, Number of run goroutines 28
Job number: 978; Counter: 972, Number of run goroutines 28
Job number: 979; Counter: 972, Number of run goroutines 28
Job number: 982; Counter: 975, Number of run goroutines 25
Job number: 981; Counter: 975, Number of run goroutines 25
Job number: 983; Counter: 975, Number of run goroutines 25
Job number: 984; Counter: 975, Number of run goroutines 25
Job number: 985; Counter: 978, Number of run goroutines 22
Job number: 987; Counter: 978, Number of run goroutines 22
Job number: 986; Counter: 978, Number of run goroutines 22
Job number: 990; Counter: 981, Number of run goroutines 19
Job number: 989; Counter: 981, Number of run goroutines 19
Job number: 988; Counter: 981, Number of run goroutines 19
Job number: 991; Counter: 984, Number of run goroutines 16
Job number: 992; Counter: 984, Number of run goroutines 16
Job number: 993; Counter: 984, Number of run goroutines 16
Job number: 995; Counter: 989, Number of run goroutines 11
Job number: 1000; Counter: 994, Number of run goroutines 6
Job number: 994; Counter: 989, Number of run goroutines 11
Job number: 996; Counter: 989, Number of run goroutines 11
Job number: 999; Counter: 994, Number of run goroutines 6
Job number: 997; Counter: 994, Number of run goroutines 6
Job number: 998; Counter: 994, Number of run goroutines 6
I've counted 1000 goroutines. 
</pre>
</details>

Запускал с десяток раз, всё время на выходе правильное количество горутин
```
I've counted 1000 goroutines.
```

Но что-то смущает в данном решении - 100 горутин в пуле параллельно работают с одной и той же внешней переменной, 
однако счетчик не сбивается. Это нормальное решение или нужно добавлять mutex.Lock/Unlock в момент обращения к переменной
(когда инкрементируем счетчик через counter++)?

**P.S. Видимо проблема кроется в том, что время вывода в STDOUT существенно больше времени работы с переменной
и поэтому не происходит одновременной записи в неё из разных горутин.
Если закомментировать строки со Sleep и Printf, то counter действительно выдает разные значения.**

### 1.2 Корректное решение (синхронизация между горутинами через канал)
Создаем буферизированный канал на 100 значений, для ограничения количества одновременно работающих воркеров:
``` golang
workers := make(chan struct{}, 100)
defer close(workers)
```
Для синхронизации между горутинами создаем небуферизированный канал:
``` golang
intCh := make(chan int)
defer close(intCh)
```
Задаем начальное значение для счетчика и отправляем его в канал:
``` golang
go func() {
	intCh <- 0
}()
```
В цикле создаем воркеры:
```golang
for i := 1; i <= 1000; i++ {
	workers <- struct{}{}
	go worker(workers, intCh, i)
}
```
Сам воркер пытается прочитать из канала, если получилось - то делает инкремент и пишет обратно в канал.
После завершения работы происходит чтение из канала pool, для возможности запуска другого воркера в нашем пуле.
```golang
func worker(pool chan struct{}, res chan int, job int) {
	defer func() {
		<-pool // Read from channel to unlock another waiting goroutines
	}()
	var result int
	for {
		select {
		case result = <-res:
			//fmt.Printf("Job number: %d; Counter: %d, Number of run goroutines %d\n", job, result, len(pool))
			result++
			res <- result
			return
		}
	}
}
```
Работает корректно и с выводом отладочной информации, и без него.

## 2. Написать программу, которая при получении в канал сигнала SIGTERM останавливается не позднее, чем за одну секунду (установить таймаут)
Объявляем функцию обработчик системных вызовов, ожидающий сигналы SIGTERM и SIGINT (для удобства отладки).
При получении сигнала запускается функция cancel, которая завершает контекст. После секундного ожидания выходим из приложения через log.Fatal:  
``` golang
func sigHandler(cancel context.CancelFunc) {
	sigChannel := make(chan os.Signal)
	signal.Notify(sigChannel, syscall.SIGTERM, syscall.SIGINT)
	for {
		select {
		case <-sigChannel:
			fmt.Printf("%s Get SIGTERM, try to stop.\n", time.Now().Format("2006/01/02 15:04:05"))
			cancel()
			time.Sleep(time.Second)
			log.Fatal("Application is killed!\n")
		}
	}
}
```

В основной программе создаем объекты воркера и контекста. Затем запускаем worker и хэндлер:  
``` golang
w := worker.NewWorker()
ctx, cancel := context.WithCancel(context.Background())

w.Run(ctx)
sigHandler(cancel)
```

Воркер каждые 100 мс запускает функцию Pinger и мониторит канал ctx.Done(), в который может быть получен сигнал завершения работы из контекста:
``` golang
func (w *Worker) Run(ctx context.Context) {
	ticker := time.NewTicker(doPeriod * time.Millisecond)

	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				resp := Pinger()
				fmt.Println(resp)
			}
		}
	}()
}
```