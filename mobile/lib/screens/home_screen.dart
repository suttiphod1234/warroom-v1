import 'package:flutter/material.dart';
import 'package:dio/dio.dart';
import 'meeting_screen.dart';

class HomeScreen extends StatefulWidget {
  final String token;
  const HomeScreen({super.key, required this.token});

  @override
  State<HomeScreen> createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  Map<String, dynamic>? _profile;

  @override
  void initState() {
    super.initState();
    _fetchProfile();
  }

  Future<void> _fetchProfile() async {
    try {
      final response = await Dio().get(
        'http://localhost:8080/api/auth/profile',
        options: Options(headers: {'Authorization': 'Bearer ${widget.token}'}),
      );
      setState(() => _profile = response.data);
    } catch (e) {
      print(e);
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: const Text('MY DASHBOARD'),
        backgroundColor: Colors.blue[900],
      ),
      body: _profile == null
          ? const Center(child: CircularProgressIndicator())
          : ListView(
              padding: const EdgeInsets.all(16),
              children: [
                _buildCard("USER", _profile!['user']['FullName'], Colors.blue),
                _buildCard("TIER", "Level ${_profile!['user']['TierLevel']}", Colors.purple),
                _buildCard("PV (Personal)", "${_profile!['wallet']['PVBalance']}", Colors.green),
                _buildCard("GV (Group)", "${_profile!['wallet']['GVBalance']}", Colors.orange),
                _buildCard("COMMISSION", "฿${_profile!['wallet']['CommissionBalance']}", Colors.amber),
                const SizedBox(height: 20),
                SizedBox(
                  height: 50,
                  child: ElevatedButton.icon(
                    onPressed: () => Navigator.push(
                      context,
                      MaterialPageRoute(builder: (context) => MeetingScreen(username: _profile!['user']['Username'])),
                    ),
                    icon: const Icon(Icons.videocam),
                    label: const Text('JOIN SECRET MEETING'),
                    style: ElevatedButton.styleFrom(backgroundColor: Colors.red[900]),
                  ),
                ),
              ],
            ),
    );
  }

  Widget _buildCard(String title, String value, Color color) {
    return Card(
      color: Colors.grey[900],
      margin: const EdgeInsets.only(bottom: 16),
      shape: RoundedRectangleBorder(
        side: BorderSide(color: color, width: 2),
        borderRadius: BorderRadius.circular(8),
      ),
      child: Padding(
        padding: const EdgeInsets.all(20),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text(title, style: TextStyle(color: Colors.grey[400], fontSize: 14)),
            const SizedBox(height: 8),
            Text(value, style: const TextStyle(color: Colors.white, fontSize: 24, fontWeight: FontWeight.bold)),
          ],
        ),
      ),
    );
  }
}
