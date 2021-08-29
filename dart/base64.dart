import 'dart:convert';

import 'dart:typed_data';

main(List<String> args) {
  //This is a base 64 string from verify me api
  var base64 ='/9j/4AAQSkZJRgABAgAAAQABAAD/ ...';
  base64 = processBase64(base64);

  //Get the Uint8List and then pass to your image Widget
  Uint8List output = Base64Decoder().convert(base64);

  // The image widget uses the Uint8List list 
  //Image.memory(output);
  print(output);
}

String processBase64(String data) {
  return data.replaceAll("\n", "");
}
