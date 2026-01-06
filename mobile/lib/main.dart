import 'package:flutter/material.dart';
import 'screens/login_screen.dart';

void main() {
  runApp(const WarRoomApp());
}

class WarRoomApp extends StatelessWidget {
  const WarRoomApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'War Room Mobile',
      theme: ThemeData(
        colorScheme: ColorScheme.fromSeed(seedColor: Colors.deepPurple),
        useMaterial3: true,
      ),
      home: const LoginScreen(),
    );
  }
}
