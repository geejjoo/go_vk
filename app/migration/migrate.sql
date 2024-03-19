
INSERT INTO Users (name, balance) VALUES
                                      ('User1', 1000),
                                      ('User2', 2000),
                                      ('User3', 3000);


INSERT INTO Quest (name, cost, times) VALUES
                                          ('Quest1', 100, "2022-11-31T09:44:00Z"),
                                          ('Quest2', 200, "2022-10-31T02:31:00Z"),
                                          ('Quest3', 300, "2022-9-31T02:00:00Z");


INSERT INTO History (questId, userId, status, timeStart, timeStop, timeDeadLine) VALUES
                                                                                     (1, 1, 'completed', "2022-9-31T02:00:00Z", "2022-9-31T04:00:00Z", "2022-9-31T05:00:00Z"),
                                                                                     (2, 2, 'completed', "2022-9-31T02:00:00Z", "2022-9-31T04:00:00Z", "2022-9-31T05:00:00Z"),
                                                                                     (3, 3, 'run', "2022-9-31T02:00:00Z", NILL, "2022-9-31T05:00:00Z");
