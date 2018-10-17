
[Source](https://medium.com/exploring-code/why-should-you-learn-go-f607681fad65 "Permalink to Why should you learn Go? – Exploring Code – Medium")

# Tại sao bạn nên học Go ? khám phá code - medium

"Go sẽ trở thành ngôn ngữ server trong tương lai."
“Go will be the server language of the future.” — Tobias Lütke, Shopify
Trong vài năm qua , đây là sự tăng trưởng mạnh mẽ của một ngôn ngữ lập trình mới : Go hay còn gọi là GoLang . Không gì có thể khiến một lập trình viên điên cuồng hơn là một ngôn ngữ mới , phải vậy không ?  Vì vậy, Tôi đã bắt đầu học Go khoảng 4 5 tháng trước và hiện tại tôi sẵn sàng để giới thiệu về việc tại sao tôi cần học Go.
Tôi không phải là đang dạy bạn đâu , Bạn có thể biết cách viết "Hello World !!" trong bài báo này .Và có nhiều bài báo online khác viết về nó . Tôi sẽ giải thích về giai đoạn hiện tại của phần mềm máy tính và tại sao chúng ta cần ngôn ngữ mo ? bởi vì chúng ta sẽ không sinh ra giải pháp nếu không có bất cứ vấn đề nào xảy ra , đúng không ?
### Giới hạn phần cứng:
Moore’s law đã thất bại.

Bộ Pentium 4 vi xử lý đầu tiên với  3.0GHz đồng hồ tốc độ được giới thiệu từ năm 2004 bởi Intel , ngày nay , Chiếc Mackbook Pro 2006 của tôi có tốc độ là 2.9Ghz. Vì vậy , trải qua gần 1 thập kỷ ,  không có quá nhiều hiệu quả trong việc xử lý thôi.BBạn có thể thấy việc so sánh tăng sức mạnh xử lý với thời gian trong biểu đồ dưới đây.

Từ biểu đồ trên bạn có thể thấy hiệu xuất của đơn luồng và tần số của bộ xử lý vẫn ổn định trong gần một thập kỷ.Nếu bạn đang nghĩ rằng việc thêm nhiều bóng bán dẫn là giải pháp, thì bạn đã sai. Điều này là do ở quy mô nhỏ hơn, một số tính chất lượng tử bắt đầu nổi lên (như đường hầm) và vì nó thực sự tốn nhiều tiền hơn để đặt nhiều bóng bán dẫn hơn (tại sao?) Và số bóng bán dẫn bạn có thể thêm vào mỗi đô la bắt đầu giảm.

Vì vậy , có một số giải pháp của vấn đề trên,

- Các nhà sản xuất bắt đầu bổ sung thêm nhiều lõi hơn cho bộ vi xử lý.
- Ngày nay, chúng tôi có CPU quad-core và octa-core.
- Chúng ta cũng đã giới thiệu về siêu phân luồng.

Đã thêm bộ nhớ cache vào bộ xử lý để tăng hiệu suất.
Nhưng các giải pháp trên cũng có những hạn chế riêng. Chúng ta không thể thêm bộ nhớ cache nhiều hơn và nhiều hơn nữa để xử lý để tăng hiệu suất như bộ nhớ cache có giới hạn vật lý: bộ nhớ cache lớn hơn, chậm hơn nó được. Thêm nhiều lõi hơn cho bộ vi xử lý cũng có chi phí của nó. Ngoài ra, điều đó không thể mở rộng đến vô thời hạn. Các bộ vi xử lý đa lõi này có thể chạy đồng thời nhiều luồng và mang lại sự đồng thời với hình ảnh. Chúng ta sẽ thảo luận điều đó sau .

Vì vậy, nếu chúng ta không thể dựa vào những cải tiến phần cứng, cách duy nhất để đi là phần mềm hiệu quả hơn để tăng hiệu suất. Nhưng thật đáng buồn, ngôn ngữ lập trình hiện đại không hiệu quả lắm.

“Bộ vi xử lý hiện đại giống như những chiếc xe hơi vui nhộn của nitro, chúng vượt trội ở phần tư dặm. Thật không may là các ngôn ngữ lập trình hiện đại giống như Monte Carlo, chúng có đầy đủ các vòng xoắn và quay. ”- David Ungar


### **Go có goroutines !!**

Như chúng ta đã thảo luận ở trên , phần cứng cần bổ xung thêm thật nhiều cores vào bộ vi xử lý để tăng hiệu xuất.Toàn bộ dữ liệu trung tâm đều chạy trên bộ vi xử lý đó và chúng mong muốn tăng số cores trong những năm tới. Thêm vào đó , các ứng dụng ngày nay sử dụng nhiều micro-services để bảo trì việc kết nối cơ sở dữ liệu ,  tin nhắn hàng đợi và lưu trữ cache. Do đó , phần mềm chúng ta phát triển và ngôn ngữ lập trình nên hỗ trợ 1 cách dễ dàng và chúng nên có thể mở rộng với với việc tăng cores.

Nhưng, một số các ngôn ngữ lập trình hiện đại (như Java , Python vân vân.) trích từ '90s môi trường đơn luồng.Hầu hết các ngôn ngữ lập trình hỗ trợ đa luồng . Nhưng vấn đề thực sự xảy ra khi thực hiện đồng thời ,  threading-locking, race conditions and deadlocks' Những điều đó làm cho nó khó khăn để tạo ra một ứng dụng đa luồng trên các ngôn ngữ đó. 

Ví dụ, tạo chuỗi mới trong Java không phải là memory efficientVì mỗi thread tiêu tốn khoảng 1MB kích thước bộ nhớ heap và cuối cùng nếu bạn bắt đầu quay hàng nghìn luồng, chúng sẽ gây áp lực to lớn trên heap và sẽ bị sập nguồn do bộ nhớ bị mất. Ngoài ra, nếu bạn muốn giao tiếp giữa hai hoặc nhiều chủ đề, nó rất khó.

Mặt khác, Go được phát hành vào năm 2009 khi bộ vi xử lý đa lõi đã có sẵn. Đó là lý do tại sao Go được xây dựng với việc lưu giữ đồng thời trong tâm trí. Go có goroutines thay vì threads. Chúng tiêu tốn gần 2KB bộ nhớ từ heap. Vì vậy, bạn có thể quay hàng triệu goroutines bất cứ lúc nào.

![][1]

Goroutines hoạt động như thế nào? Reffrance: http://golangtutorials.blogspot.in/2011/06/goroutines.html

**Các lợi ích khác là:**

* Goroutines có ngăn xếp phân đoạn có thể phát triển. Điều đó có nghĩa là họ sẽ chỉ sử dụng nhiều bộ nhớ hơn khi cần thiết.
* Goroutines có thời gian khởi động nhanh hơn threads.
* Goroutines đi kèm với các phần tử tích hợp sẵn để giao tiếp an toàn giữa chúng.
* Goroutines cho phép bạn tránh phải sử dụng khóa mutex khi chia sẻ cấu trúc dữ liệu.
* Ngoài ra, các luồng goroutines và OS không có ánh xạ 1: 1. Một goroutine đơn có thể chạy trên nhiều luồng. Goroutine được ghép thành một số lượng nhỏ các chuỗi hệ điều hành.

Tất cả các điểm trên, làm cho Go rất mạnh mẽ để xử lý đồng thời như Java, C và C ++ trong khi vẫn giữ đoạn mã thực hiện đồng thời và đẹp như Erlang.

Go chạy trực tiếp trên phần cứng cơ bản.

Một lợi ích đáng kể nhất của việc sử dụng C, C ++ so với các ngôn ngữ bậc cao hiện đại khác như Java / Python là hiệu năng của chúng. Bởi vì C / C ++ được biên dịch và không được giải thích.

Bộ xử lý hiểu các tệp nhị phân. Nói chung, khi bạn xây dựng một ứng dụng bằng cách sử dụng Java hoặc các ngôn ngữ dựa trên JVM khác khi biên dịch dự án của bạn, nó biên dịch mã có thể đọc được thành mã byte có thể được JVM hoặc các máy ảo khác chạy trên hệ điều hành cơ bản. Trong khi thực hiện, VM diễn giải các bytecode đó và chuyển đổi chúng thành các tập tin nhị phân mà bộ xử lý có thể hiểu được.

Các bước thực hiện cho các ngôn ngữ dựa trên máy ảo
Trong khi ở phía bên kia, C / C ++ không thực hiện trên máy ảo và loại bỏ một bước từ chu kỳ thực hiện và tăng hiệu suất. Nó trực tiếp biên dịch mã có thể đọc được của con người thành mã nhị phân.

Nhưng, giải phóng và phân bổ biến trong những ngôn ngữ đó là một nỗi đau lớn. Trong khi hầu hết các ngôn ngữ lập trình xử lý phân bổ đối tượng và loại bỏ bằng cách sử dụng Garbage Collector hoặc thuật toán đếm tham chiếu.

Đi mang lại tốt nhất của cả hai thế giới. Giống như các ngôn ngữ cấp thấp hơn như C / C ++, Go là ngôn ngữ được biên dịch. Điều đó có nghĩa là hiệu suất gần như gần hơn với các ngôn ngữ cấp thấp hơn. Nó cũng sử dụng bộ thu gom rác để phân bổ và loại bỏ đối tượng. Vì vậy, không có thêm malloc() và free() statements!!! cool!!!

Mã được viết bằng Go rất dễ bảo trì.
Hãy để tôi nói với bạn một điều. Go không có cú pháp lập trình điên như các ngôn ngữ khác. Nó có cú pháp rất gọn gàng và sạch sẽ.

Các nhà thiết kế của Go tại google đã có điều này trong tâm trí khi họ đã tạo ra ngôn ngữ. Vì google có cơ sở mã rất lớn và hàng nghìn nhà phát triển đang làm việc trên cùng một mã cơ sở đó nên mã đơn giản phải hiểu cho các nhà phát triển khác và một đoạn mã nên có tác dụng phụ tối thiểu trên một đoạn mã khác. Điều đó sẽ làm cho mã dễ bảo trì và dễ sửa đổi.
## Go Cố tình bỏ qua nhiều tính năng của các ngôn ngữ OOP hiện đại.
Không có class nào. Mọi thứ chỉ được chia thành các gói. Go chỉ có cấu trúc thay vì các class.
Không hỗ trợ kế thừa. Điều đó sẽ làm cho mã dễ sửa đổi. Trong các ngôn ngữ khác như Java / Python, nếu lớp ABC kế thừa lớp XYZ và bạn thực hiện một số thay đổi trong lớp XYZ, thì có thể tạo ra một số tác dụng phụ trong các lớp khác kế thừa XYZ. Bằng cách loại bỏ kế thừa, Go làm cho nó dễ hiểu mã hơn (vì không có siêu class để xem xét trong khi nhìn vào một đoạn mã).

Không có hàm tạo.
Không có chú thích.
Không có generics.
Không có ngoại lệ.

Những thay đổi trên làm cho Go rất khác với các ngôn ngữ khác và nó làm cho lập trình trong Go khác với các ngôn ngữ khác. Bạn có thể không thích một số điểm từ trên cao. Tuy nhiên, nó không giống như bạn không thể code ứng dụng của bạn mà không có các tính năng trên. Tất cả những gì bạn phải làm là viết 2-3 dòng nữa. Nhưng về mặt tích cực, nó sẽ làm cho mã của bạn sạch hơn và thêm rõ ràng hơn cho mã của bạn.

Khả năng đọc mã vs, Hiệu quả.
Biểu đồ trên hiển thị rằng Go gần như hiệu quả như C / C ++, trong khi vẫn giữ cú pháp mã đơn giản như Ruby, Python và các ngôn ngữ khác. Đó là một tình huống có lợi cho cả con người và bộ vi xử lý !!!

Không giống như các ngôn ngữ mới khác như Swift, cú pháp của Go rất ổn định. Nó vẫn như cũ kể từ khi phát hành công khai ban đầu 1.0, trở lại trong năm 2012. Điều đó làm cho nó tương thích ngược.

## Go được hỗ trợ bởi Google.
Tôi biết đây không phải là một lợi thế kỹ thuật trực tiếp. Tuy nhiên, Go được thiết kế và hỗ trợ bởi Google. Google có một trong những cơ sở hạ tầng đám mây lớn nhất trên thế giới và nó được mở rộng quy mô. Go được thiết kế bởi Google để giải quyết các vấn đề của họ về hỗ trợ khả năng mở rộng và hiệu quả. Đó là những vấn đề tương tự bạn sẽ phải đối mặt trong khi tạo ra các máy chủ của riêng bạn.

Thêm vào đó Go cũng được sử dụng bởi một số công ty lớn như Adobe, BBC, IBM, Intel và thậm chí Trung bình (Nguồn: https://github.com/golang/go/wiki/GoUsers)

Kết luận:
Mặc dù Go rất khác với các ngôn ngữ hướng đối tượng khác, nó vẫn là cùng một con thú. Go cung cấp cho bạn hiệu suất cao như C / C ++, xử lý đồng thời siêu hiệu quả như Java và thú vị với mã như Python / Perl.

Nếu bạn không có bất kỳ kế hoạch học Go, tôi vẫn sẽ nói giới hạn phần cứng đặt áp lực cho chúng tôi, các nhà phát triển phần mềm để viết mã siêu hiệu quả. Nhà phát triển cần hiểu phần cứng và làm cho chương trình của họ tối ưu hóa phù hợp. Phần mềm được tối ưu hóa có thể chạy trên phần cứng rẻ hơn và chậm hơn (như thiết bị IOT) và tác động tổng thể tốt hơn đến trải nghiệm người dùng cuối.

~ Nếu bạn thích bài viết, hãy nhấp vào 💚 bên dưới để nhiều người hơn có thể nhìn thấy nó! Ngoài ra, bạn có thể theo dõi tôi trên Trung bình hoặc trên Blog của tôi, để bạn nhận được thông tin cập nhật về các bài viết trong tương lai trên Go !! ~

![][3]

Go takes good of both the worlds. Easy to write concurrent and efficient to manage concurrency

[1]: https://cdn-images-1.medium.com/max/1600/1*NFojvbkdRkxz0ZDbu4ysNA.jpeg
[2]: https://blog.golang.org/concurrency-is-not-parallelism
[3]: https://cdn-images-1.medium.com/max/1600/1*xbsHBQJReC5l_VO4XgNSIQ.png

  
