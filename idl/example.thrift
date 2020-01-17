namespace go example

struct Data {
    1:string text
}


service exampleService {
    Data do_format(1:Data data)
}