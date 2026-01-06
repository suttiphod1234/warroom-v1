import 'package:flutter/material.dart';
import 'package:omni_jitsi_meet/omni_jitsi_meet.dart';

class MeetingScreen extends StatefulWidget {
  final String username;
  const MeetingScreen({super.key, required this.username});

  @override
  State<MeetingScreen> createState() => _MeetingScreenState();
}

class _MeetingScreenState extends State<MeetingScreen> {
  final _roomController = TextEditingController(text: "WarRoom-Secret-HQ-001");

  Future<void> _joinMeeting() async {
    try {
      var options = JitsiMeetingOptions(
        roomNameOrUrl: _roomController.text,
        serverUrl: "https://meet.jit.si",
        subject: "War Room Briefing",
        userDisplayName: widget.username,
        userEmail: "${widget.username}@warroom.com",
        userAvatarUrl: "", // Can be added later
        isAudioMuted: true,
        isVideoMuted: true,
      );

      await JitsiMeet.joinMeeting(options);
    } catch (error) {
      debugPrint("error: $error");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black,
      appBar: AppBar(
        title: const Text('SECURE CHANNEL'),
        backgroundColor: Colors.red[900],
      ),
      body: Padding(
        padding: const EdgeInsets.all(24.0),
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Icon(Icons.videocam_off, size: 80, color: Colors.red),
            const SizedBox(height: 20),
            const Text(
              'SECRET MEETING ROOM',
              style: TextStyle(fontSize: 24, fontWeight: FontWeight.bold, color: Colors.white),
            ),
            const SizedBox(height: 10),
            const Text(
              'End-to-end encrypted video channel.',
              style: TextStyle(color: Colors.grey),
            ),
            const SizedBox(height: 40),
            TextField(
              controller: _roomController,
              style: const TextStyle(color: Colors.white),
              decoration: const InputDecoration(
                labelText: 'Room ID',
                labelStyle: TextStyle(color: Colors.grey),
                enabledBorder: OutlineInputBorder(borderSide: BorderSide(color: Colors.red)),
                focusedBorder: OutlineInputBorder(borderSide: BorderSide(color: Colors.redAccent)),
              ),
            ),
            const SizedBox(height: 30),
            SizedBox(
              width: double.infinity,
              height: 50,
              child: ElevatedButton.icon(
                onPressed: _joinMeeting,
                icon: const Icon(Icons.login),
                label: const Text('JOIN SECURE MEETING'),
                style: ElevatedButton.styleFrom(backgroundColor: Colors.red[700]),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
