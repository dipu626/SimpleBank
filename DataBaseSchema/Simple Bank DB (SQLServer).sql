CREATE TABLE [Account] (
  [Id] int PRIMARY KEY IDENTITY(1, 1),
  [FirstName] nvarchar(255) NOT NULL,
  [Ballance] bigint NOT NULL,
  [Currency] nvarchar(255) NOT NULL,
  [CreatedAt] timestamptz NOT NULL DEFAULT 'now()'
)
GO

CREATE TABLE [Entries] (
  [Id] int PRIMARY KEY IDENTITY(1, 1),
  [AccountNumber] bigint,
  [Amount] bigint NOT NULL,
  [CreatedAt] timestamptz NOT NULL DEFAULT 'now()'
)
GO

CREATE TABLE [Transfers] (
  [Id] bigint PRIMARY KEY IDENTITY(1, 1),
  [FromAccount] int NOT NULL,
  [ToAccount] int NOT NULL,
  [Amount] bigint NOT NULL,
  [CreatedAt] timestamptz NOT NULL DEFAULT 'now()'
)
GO

CREATE INDEX [Account_index_0] ON [Account] ("Id")
GO

CREATE INDEX [Entries_index_1] ON [Entries] ("AccountNumber")
GO

CREATE INDEX [Transfers_index_2] ON [Transfers] ("FromAccount")
GO

CREATE INDEX [Transfers_index_3] ON [Transfers] ("ToAccount")
GO

CREATE INDEX [Transfers_index_4] ON [Transfers] ("FromAccount", "ToAccount")
GO

EXEC sp_addextendedproperty
@name = N'Column_Description',
@value = 'Can be negative or positive',
@level0type = N'Schema', @level0name = 'dbo',
@level1type = N'Table',  @level1name = 'Entries',
@level2type = N'Column', @level2name = 'Amount';
GO

EXEC sp_addextendedproperty
@name = N'Column_Description',
@value = 'Must be positive',
@level0type = N'Schema', @level0name = 'dbo',
@level1type = N'Table',  @level1name = 'Transfers',
@level2type = N'Column', @level2name = 'Amount';
GO

ALTER TABLE [Entries] ADD FOREIGN KEY ([AccountNumber]) REFERENCES [Account] ([Id])
GO

ALTER TABLE [Transfers] ADD FOREIGN KEY ([FromAccount]) REFERENCES [Account] ([Id])
GO

ALTER TABLE [Transfers] ADD FOREIGN KEY ([ToAccount]) REFERENCES [Account] ([Id])
GO
